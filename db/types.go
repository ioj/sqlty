package db

import (
	"context"
	"fmt"

	"github.com/ioj/sqlty/stmt"
	"github.com/jackc/pgx/v4"
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
}

type pgTypes struct {
	types map[uint32]*pgType
	enums map[uint32][]string
}

func newPgTypes(ctx context.Context, db *pgx.Conn) (*pgTypes, error) {
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

	for rows.Next() {
		t := &pgType{}
		if err := rows.Scan(&t.OID, &t.Namespace, &t.Name, &t.Type,
			&t.Category, &t.Elem, &t.NotNull, &t.BaseType); err != nil {
			return nil, err
		}

		pt.types[t.OID] = t
	}

	if err := rows.Err(); err != nil {
		return nil, err
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

	return pt, nil
}

func (t *pgType) Fullname() string {
	return t.Namespace + "." + t.Name
}

func (pt *pgTypes) Type(oid uint32, notnull bool) (*stmt.Type, error) {
	pgtype, ok := pt.types[oid]
	if !ok {
		return nil, fmt.Errorf("no type mapping for OID = %v", oid)
	}

	// Handle array types
	if pgtype.Category == 'A' {
		if pgtype.Elem == 0 {
			return nil, fmt.Errorf("array type with typelem = 0, it shouldn't happen")
		}

		pgtype, err := pt.Type(pgtype.Elem, notnull)
		if err != nil {
			return nil, err
		}

		pgtype.ZeroValue = fmt.Sprintf("make([]%v)", pgtype.Name)
		pgtype.Name = "[]" + pgtype.Name
		pgtype.Nullable = true

		return pgtype, nil
	}

	type T struct {
		fullname string
		notnull  bool
	}

	t := T{pgtype.Fullname(), notnull}

	switch t {
	case T{"pg_catalog.int2", true}:
		return &stmt.Type{Name: "int", ZeroValue: "0", Nullable: false}, nil
	case T{"pg_catalog.int2", false}:
		return &stmt.Type{Name: "*int", ZeroValue: "new(int)", Nullable: true}, nil
	case T{"pg_catalog.int8", true}:
		return &stmt.Type{Name: "int", ZeroValue: "0", Nullable: false}, nil
	case T{"pg_catalog.int8", false}:
		return &stmt.Type{Name: "*int", ZeroValue: "new(int)", Nullable: true}, nil
	case T{"pg_catalog.text", true}:
		return &stmt.Type{Name: "string", ZeroValue: "\"\"", Nullable: false}, nil
	case T{"pg_catalog.text", false}:
		return &stmt.Type{Name: "*string", ZeroValue: "new(string)", Nullable: true}, nil
	}

	return nil, fmt.Errorf("unknown type oid = %v, notnull = %v", oid, notnull)
}
