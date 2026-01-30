package db

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"github.com/ioj/sqlty/helpers"
	"github.com/ioj/sqlty/stmt"
	"github.com/jackc/pgx/v5"
	"github.com/serenize/snaker"
)

var (
	errVoid      = errors.New("pg_catalog.void type")
	errComposite = errors.New("composite type")
)

type pgType struct {
	// Row identifier
	OID uint32

	// Name of the namespace
	Namespace string

	// Data type name
	Name string

	// b - a base type
	// c - a composite type (e.g., a table's row type)
	// d - a domain
	// e - an enum type
	// p - pseudo-type
	// r - a range type
	Type rune

	// See https://www.postgresql.org/docs/current/catalog-pg-type.html
	// A - array
	// E - enum
	// and a few others
	Category rune

	// If not zero, then an OID of the underlying base type for arrays
	Elem uint32

	// Not-null constraint for domain types
	NotNull bool

	// If this is a domain, then typbasetype identifies the type that
	// this one is based on. Zero if this type is not a domain.
	BaseType uint32

	// If this is a composite type, RelID is a pg_attribute reference
	// for fields list
	RelID uint32

	// If UniqueName is true, there are no other types with the same name,
	// but different namespace.
	UniqueName bool

	// Used is true for composite types that are required for generated queries.
	Used bool

	// List of underlying OIDs for composite types.
	Fields []uint32
}

type pgTypes struct {
	db *pgx.Conn

	types map[uint32]*pgType
	enums map[uint32][]string

	translations map[PGTypeDef]stmt.Type
}

type PGTypeDef struct {
	Fullname string `yaml:"name"`
	NotNull  bool   `yaml:"notNull"`
}

type PGTypeTranslation struct {
	Fullname string    `yaml:"name" mapstructure:"name"`
	NotNull  bool      `yaml:"notNull"`
	To       stmt.Type `yaml:"to"`
}

type PGTypeTranslationsFile struct {
	Types []PGTypeTranslation `yaml:"types"`
}

func newPgTypes(ctx context.Context, db *pgx.Conn, types []PGTypeTranslation) (*pgTypes, error) {
	pt := &pgTypes{
		db:    db,
		types: make(map[uint32]*pgType),
		enums: make(map[uint32][]string),
	}

	// Load all types from the database

	rows, err := db.Query(ctx, `
		SELECT
			t.oid, n.nspname, t.typname, t.typtype, t.typcategory,
			t.typelem, t.typnotnull, t.typbasetype, t.typrelid
		FROM pg_type t
		LEFT JOIN pg_namespace n ON n.oid = t.typnamespace
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	uniqueNames := make(map[string][]uint32)

	for rows.Next() {
		t := &pgType{UniqueName: true}
		if err := rows.Scan(&t.OID, &t.Namespace, &t.Name, &t.Type,
			&t.Category, &t.Elem, &t.NotNull, &t.BaseType, &t.RelID); err != nil {
			return nil, err
		}

		n := snaker.SnakeToCamel(t.Name)

		// PostgreSQL defines type names of arrays of type as _typename.
		// SnakeToCamel removes the underscore prefix. Let's prevent that.
		if t.Name[0] == '_' {
			n = "_" + n
		}

		uniqueNames[n] = append(uniqueNames[n], t.OID)

		pt.types[t.OID] = t
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Fix UniqueName in the map
	for _, oids := range uniqueNames {
		if len(oids) > 1 {
			for _, oid := range oids {
				pt.types[oid].UniqueName = false
			}
		}
	}

	// Load all enum labels from the database

	rows, err = db.Query(ctx, `
		SELECT enumtypid, enumlabel
		FROM pg_enum
		ORDER BY enumtypid, enumsortorder
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var typoid uint32
		var label string

		if err := rows.Scan(&typoid, &label); err != nil {
			return nil, err
		}

		pt.enums[typoid] = append(pt.enums[typoid], label)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	pt.translations = make(map[PGTypeDef]stmt.Type)
	for _, t := range types {
		ptd := PGTypeDef{Fullname: t.Fullname, NotNull: t.NotNull}
		pt.translations[ptd] = t.To
	}

	return pt, nil
}

func (t *pgType) Fullname() string {
	return t.Namespace + "." + t.Name
}

func (t *pgType) GolangName() string {
	if t.UniqueName || t.Namespace == "public" {
		return snaker.SnakeToCamel(t.Name)
	}

	return snaker.SnakeToCamel(t.Namespace + "_" + t.Name)
}

func (pt *pgTypes) Enums() []*stmt.Enum {
	var enums []*stmt.Enum

	for _, t := range pt.types {
		if t.Type != 'e' {
			continue
		}

		e := &stmt.Enum{Name: t.GolangName(), Values: pt.enums[t.OID]}
		enums = append(enums, e)
	}

	sort.Slice(enums, func(i, j int) bool {
		return enums[i].Name < enums[j].Name
	})

	return enums
}

func (pt *pgTypes) CompositeTypes(ctx context.Context) ([]*stmt.Struct, error) {
	var types []*stmt.Struct

	for _, t := range pt.types {
		if t.Type != 'c' {
			continue
		}

		if !t.Used {
			continue
		}

		composite, err := pt.CompositeFields(ctx, t.OID)
		if err != nil {
			return nil, err
		}

		types = append(types, composite)
	}

	sort.Slice(types, func(i, j int) bool {
		return types[i].Name < types[j].Name
	})

	return types, nil
}

func (pt *pgTypes) CompositeFields(ctx context.Context, oid uint32) (*stmt.Struct, error) {
	pgtype, ok := pt.types[oid]
	if !ok {
		return nil, fmt.Errorf("no type mapping for OID = %v", oid)
	}

	if pgtype.Type != 'c' {
		return nil, fmt.Errorf("not a composite type: %v", pgtype.Type)
	}

	if pgtype.RelID == 0 {
		return nil, fmt.Errorf("composite type with relid = 0 (%v)", oid)
	}

	pt.types[oid].Used = true

	// DontRender is true, because composite types are rendered in a separate file.
	s := &stmt.Struct{Name: pgtype.GolangName(), IsCompositeType: true}

	rows, err := pt.db.Query(ctx, `
		SELECT atttypid, attname, attnotnull
		FROM pg_attribute
		WHERE attrelid = $1
		ORDER BY attnum`, pgtype.RelID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	normalizer := helpers.NewStructFieldNormalizer()
	for rows.Next() {
		var oid uint32
		var name string
		var notnull bool

		if err := rows.Scan(&oid, &name, &notnull); err != nil {
			return nil, err
		}

		gotype, err := pt.Type(oid, notnull)
		if err != nil {
			return nil, err
		}

		nname, err := normalizer.Add(name, false)
		if err != nil {
			return nil, err
		}

		s.Params = append(s.Params, stmt.Param{Name: nname, Type: *gotype})
	}

	return s, nil
}

func (pt *pgTypes) Type(oid uint32, notnull bool) (*stmt.Type, error) {
	pgtype, ok := pt.types[oid]
	if !ok {
		return nil, fmt.Errorf("no type mapping for OID = %v", oid)
	}

	t := PGTypeDef{Fullname: pgtype.Fullname(), NotNull: notnull || pgtype.NotNull}

	stmttype, ok := pt.translations[t]
	if ok {
		return &stmttype, nil
	}

	if t.NotNull {
		// nullable types can deserialize non-nullable ones
		t2 := PGTypeDef{Fullname: pgtype.Fullname()}
		stmttype, ok := pt.translations[t2]
		if ok {
			return &stmttype, nil
		}
	}

	if pgtype.Type == 'e' {
		goname := pgtype.GolangName()
		if notnull {
			return &stmt.Type{Name: goname, ZeroValue: fmt.Sprintf("%v(\"\")", goname), Nullable: false}, nil
		}
		return &stmt.Type{Name: "*" + pgtype.GolangName(), ZeroValue: "new(" + pgtype.GolangName() + ")", Nullable: true}, nil
	}

	if t.Fullname == "pg_catalog.void" {
		return nil, errVoid
	}

	if pgtype.Type == 'c' {
		return nil, errComposite
	}

	return nil, fmt.Errorf("unknown type oid = %v, name = %v, notnull = %v", oid, t.Fullname, t.NotNull)
}
