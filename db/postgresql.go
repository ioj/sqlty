package db

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ioj/sqlty/stmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type Resolver struct {
	ctx  context.Context
	conn *pgconn.PgConn
	db   *pgx.Conn

	types *pgTypes
}

type pgAttr struct {
	attrelid uint32
	attnum   uint16
	notnull  bool
}

func NewResolver(ctx context.Context, connString string, types []PGTypeTranslation) (*Resolver, error) {
	var err error

	r := &Resolver{ctx: ctx}
	r.conn, err = pgconn.Connect(r.ctx, connString)
	if err != nil {
		return nil, err
	}

	r.db, err = pgx.Connect(r.ctx, connString)
	if err != nil {
		return nil, err
	}

	r.types, err = newPgTypes(ctx, r.db, types)
	return r, err
}

func (r *Resolver) Close() error {
	dberr := r.db.Close(r.ctx)
	connerr := r.conn.Close(r.ctx)

	if dberr != nil {
		return dberr
	}

	return connerr
}

func (r *Resolver) Enums() []*stmt.Enum {
	return r.types.Enums()
}

func (r *Resolver) CompositeTypes(ctx context.Context) ([]*stmt.Struct, error) {
	return r.types.CompositeTypes(ctx)
}

func (r *Resolver) getNullableAttrs(ctx context.Context, attrs []*pgAttr) error {
	if len(attrs) == 0 {
		return nil
	}

	var params []string

	for _, a := range attrs {
		params = append(params, fmt.Sprintf("(%d,%d)", a.attrelid, a.attnum))
	}

	stmt := fmt.Sprintf(`
		SELECT attrelid, attnum
		FROM pg_attribute WHERE attnotnull = true
			AND (attrelid, attnum) IN (%v)
		`, strings.Join(params, ","))

	rows, err := r.db.Query(ctx, stmt)
	if err != nil {
		return err
	}
	defer rows.Close()

	notnullmap := make(map[pgAttr]bool)
	for rows.Next() {
		a := pgAttr{}
		if err := rows.Scan(&a.attrelid, &a.attnum); err != nil {
			return err
		}
		notnullmap[a] = true
	}

	for _, a := range attrs {
		if _, ok := notnullmap[*a]; ok {
			a.notnull = true
		}
	}

	return nil
}

func (r *Resolver) ResolveTypes(ctx context.Context, query string, notnulls []bool) ([]stmt.Type, *stmt.Struct, error) {
	res, err := r.conn.Prepare(ctx, "", query, nil)
	if err != nil {
		return nil, nil, err
	}

	if len(res.ParamOIDs) != len(notnulls) {
		return nil, nil, fmt.Errorf("invalid number of parameters. expected %v, got %v",
			len(notnulls), len(res.ParamOIDs))
	}

	var params []stmt.Type
	for idx, p := range res.ParamOIDs {
		gotype, err := r.types.Type(p, notnulls[idx])
		if err != nil {
			return nil, nil, err
		}

		if gotype == nil {
			return nil, nil, fmt.Errorf("no type for %v", p)
		}

		params = append(params, *gotype)
	}

	attrs := make([]*pgAttr, len(res.Fields))

	for n, f := range res.Fields {
		attrs[n] = &pgAttr{attrelid: f.TableOID, attnum: f.TableAttributeNumber}
	}

	if err := r.getNullableAttrs(ctx, attrs); err != nil {
		return nil, nil, err
	}

	if len(res.Fields) == 1 {
		// If there's only one return field, there's a possibility that it's one of two supported
		// edge cases -- a function returning either void or a composite type (e.g. table row).
		// Handle it here.
		f := res.Fields[0]
		gotype, err := r.types.Type(f.DataTypeOID, attrs[0].notnull)
		switch {
		case err == nil:
			// It's just a normal return type.
			return params, &stmt.Struct{Params: []stmt.Param{{Name: string(f.Name), Type: *gotype}}}, nil
		case errors.Is(err, errVoid):
			// Query returns void.
			return params, nil, nil
		case errors.Is(err, errComposite):
			// Query returns a composite type.
			returns, err := r.types.CompositeFields(ctx, f.DataTypeOID)
			if err != nil {
				return nil, nil, err
			}
			return params, returns, err
		default:
			return nil, nil, err
		}
	}

	// For 2+ return fields, iterate through them and build a list of their Go equivalents.
	returns := &stmt.Struct{}
	for n, f := range res.Fields {
		gotype, err := r.types.Type(f.DataTypeOID, attrs[n].notnull)
		if err != nil {
			return nil, nil, err
		}

		p := stmt.Param{Name: string(f.Name), Type: *gotype}
		returns.Params = append(returns.Params, p)
	}

	return params, returns, nil
}
