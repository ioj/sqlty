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

func (r *Resolver) ResolveTypes(ctx context.Context, query string, notnulls []bool) ([]stmt.Type, []stmt.Param, error) {
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

	var returnvals []stmt.Param
	for n, f := range res.Fields {
		p := stmt.Param{Name: string(f.Name)}
		gotype, err := r.types.Type(f.DataTypeOID, attrs[n].notnull)
		switch {
		case err == errVoid:
			// A special case for functions that return void. As there is no
			// void type in go, we simply ignore this return type.
			if len(res.Fields) == 1 {
				return params, nil, nil
			}

			// This probably needs to be fixed if it can happen.
			return nil, nil, errors.New("more than one return value including void. don't know what to do")
		case err != nil:
			return nil, nil, err
		}

		p.Type = *gotype
		returnvals = append(returnvals, p)
	}

	return params, returnvals, nil
}
