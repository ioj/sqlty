package db

import (
	"context"
	"errors"
	"fmt"
	"sort"

	"github.com/ioj/sqlty/stmt"
	"github.com/jackc/pgx/v4"
	"github.com/serenize/snaker"
)

var errVoid = errors.New("pg_catalog.void type")

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

	// If UniqueName is true, there are no other types with the same name,
	// but different namespace.
	UniqueName bool
}

type pgTypes struct {
	types map[uint32]*pgType
	enums map[uint32][]string

	translations map[PGTypeDef]stmt.Type
}

type PGTypeDef struct {
	Fullname string `yaml:"name"`
	NotNull  bool   `yaml:"notNull"`
}

type PGTypeTranslation struct {
	Fullname string    `yaml:"name"`
	NotNull  bool      `yaml:"notNull"`
	To       stmt.Type `yaml:"to"`
}

type PGTypeTranslationsFile struct {
	Types []PGTypeTranslation `yaml:"types"`
}

func newPgTypes(ctx context.Context, db *pgx.Conn, types []PGTypeTranslation) (*pgTypes, error) {
	pt := &pgTypes{
		types: make(map[uint32]*pgType),
		enums: make(map[uint32][]string),
	}

	// Load all types from the database

	rows, err := db.Query(ctx, `
		SELECT
			t.oid, n.nspname, t.typname, t.typtype, t.typcategory,
			t.typelem, t.typnotnull, t.typbasetype
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
			&t.Category, &t.Elem, &t.NotNull, &t.BaseType); err != nil {
			return nil, err
		}

		n := snaker.SnakeToCamel(t.Name)
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

	return nil, fmt.Errorf("unknown type oid = %v, name = %v, notnull = %v", oid, t.Fullname, t.NotNull)
}
