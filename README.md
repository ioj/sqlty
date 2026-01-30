# SQLty

**Type-safe SQL for Go.** Write raw SQL, get fully typed Go functions.

[![Go Reference](https://pkg.go.dev/badge/github.com/ioj/sqlty.svg)](https://pkg.go.dev/github.com/ioj/sqlty)
[![Go Report Card](https://goreportcard.com/badge/github.com/ioj/sqlty)](https://goreportcard.com/report/github.com/ioj/sqlty)

---

## Why SQLty?

```sql
/* @name GetUserByEmail @one */
SELECT id, email, created_at FROM users WHERE email = :email;
```

⬇️ generates

```go
func (q *Queries) GetUserByEmail(ctx context.Context, email string) (*GetUserByEmailRow, error)

type GetUserByEmailRow struct {
    ID        int64
    Email     string
    CreatedAt time.Time
}
```

**No ORMs. No schema mapping. No runtime reflection.** SQLty connects to your PostgreSQL database, resolves types from the actual schema, and generates compile-time safe Go code.

---

## Features

- **Type-safe queries** — Parameters and return types resolved from your live database schema
- **Multiple execution modes** — `@one`, `@many`, `@exec`, and `@cursor` for streaming
- **PostgreSQL native** — Full support for enums, composite types, arrays, and JSON
- **Spread parameters** — Expand slices into `IN` clauses with `:ids...`
- **Bulk inserts** — Struct spread for inserting multiple rows in one query
- **Incremental builds** — Only regenerate changed queries

---

## Quick Start

**Install:**

```bash
go install github.com/ioj/sqlty@latest
```

**Configure** `sqlty.yaml`:

```yaml
dburl: postgres://user:pass@localhost:5432/mydb
paths:
  source: ./queries
  output: ./db
package: db
```

**Write SQL** in `queries/users.sql`:

```sql
/* @name ListActiveUsers @many */
SELECT id, email, name
FROM users
WHERE active = true
ORDER BY created_at DESC;

/* @name CreateUser @exec */
INSERT INTO users (email, name) VALUES (:email, :name);

/* @name GetUsersByIDs @many */
SELECT * FROM users WHERE id IN (:ids...);
```

**Generate:**

```bash
sqlty
```

**Use:**

```go
package main

import (
    "context"
    "yourproject/db"
    "github.com/jackc/pgx/v5"
)

func main() {
    conn, _ := pgx.Connect(context.Background(), connString)
    q := db.New(conn)

    // Fully typed - IDE autocomplete works
    users, _ := q.ListActiveUsers(ctx)

    // Spread parameters expand automatically
    users, _ = q.GetUsersByIDs(ctx, []int64{1, 2, 3})

    // Transactions via context
    tx, _ := conn.Begin(ctx)
    ctx = context.WithValue(ctx, db.CtxDBKey, tx)
    q.CreateUser(ctx, "alice@example.com", "Alice")
    tx.Commit(ctx)
}
```

---

## Annotation Reference

| Annotation | Description |
|------------|-------------|
| `@name Name` | Generated function name |
| `@one` | Returns single row |
| `@many` | Returns slice of rows |
| `@exec` | No return value |
| `@template cursor` | Streaming cursor for large result sets |
| `@param p -> type` | Override parameter type (e.g., `scalar`, `spread`) |
| `@notNull p` | Mark nullable column as non-null |
| `@return Name` | Custom return struct name |

---

## Generated Files

```
db/
├── get_user_by_email.gen.go    # Query functions
├── list_active_users.gen.go
├── enums.sqlty.gen.go          # PostgreSQL enums → Go types
├── composite_types.sqlty.gen.go # Row/composite types
└── db.sqlty.gen.go             # Queries struct & helpers
```

---

## Requirements

- Go 1.21+
- PostgreSQL
- pgx/v5

---

## License

MIT
