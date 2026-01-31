# SQLty

**SQL. Typed. Thank You.**

[![Go Reference](https://pkg.go.dev/badge/github.com/ioj/sqlty.svg)](https://pkg.go.dev/github.com/ioj/sqlty)
[![Go Report Card](https://goreportcard.com/badge/github.com/ioj/sqlty)](https://goreportcard.com/report/github.com/ioj/sqlty)

---

## The Pitch

You know that feeling when you're 3 beers deep, deploying to prod on a Friday,
and you realize your hand-crafted SQL has a typo in a column name that only
manifests when `user_type = 'enterprise'`?

No? Just me? Cool.

Anyway, I wrote SQLty.

```sql
/* @name GetUserByEmail @one */
SELECT id, emial, created_at FROM users WHERE email = :email;
--         ^^^^^ skill issue detected at compile time, not 3am
```

SQLty connects to your _actual production database_ (or staging, if you're a
coward), inspects the schema, and generates Go code that refuses to compile when
you inevitably fat-finger something.

---

## What Is This

You write SQL. Actual SQL. The kind your DBA respects.

```sql
/* @name GetUserByEmail @one */
SELECT id, email, created_at FROM users WHERE email = :email;
```

SQLty turns it into:

```go
func (q *Queries) GetUserByEmail(ctx context.Context, email string) (*GetUserByEmailRow, error)

type GetUserByEmailRow struct {
    ID        int64
    Email     string
    CreatedAt time.Time
}
```

That's it. That's the tweet.

---

## Why Not Just Use [ORM]?

Because `user.Preload("Posts.Comments.Author.Profile.Settings").Find(&users)` is
not a query plan, it's a cry for help.

Because I trust my database more than I trust an abstraction layer written by
someone who thinks N+1 is a synth-pop band.

Because I want to write:

```sql
SELECT u.*, COUNT(p.id) as post_count
FROM users u
LEFT JOIN posts p ON p.user_id = u.id
WHERE u.created_at > NOW() - INTERVAL '30 days'
GROUP BY u.id
HAVING COUNT(p.id) > 5
ORDER BY post_count DESC
```

And not whatever Rube Goldberg method chaining produces the same thing in your
ORM of choice this week.

---

## Features (The Serious Section Your Manager Can Read)

- **Type-safe queries** — Parameters and return types resolved from your live
  database schema
- **Multiple execution modes** — `@one`, `@many`, `@exec`, and `@cursor` for
  when you're feeling fancy
- **PostgreSQL native** — Full support for enums, composite types, arrays, and
  JSON(B)
- **Spread parameters** — `WHERE id IN (:ids...)` expands slices automatically
- **Bulk inserts** — Shove entire structs into INSERT statements
- **Incremental builds** — Only regenerates what changed (we respect your CPU)
- **Minimal dependencies** — Just pgx/v5 and the standard library. No dependency
  trees that make your security team cry

---

## Quick Start

**Install:**

```bash
go install github.com/ioj/sqlty@latest
```

**Configure** `sqlty.yaml`:

```yaml
dburl: postgres://user:pass@localhost:5432/mydb # your production creds here (jk) (unless?)
paths:
  source: ./queries
  output: ./db
package: db
```

**Write SQL** in `queries/`:

```sql
/* @name ListActiveUsers @many */
SELECT id, email, name
FROM users
WHERE active = true
ORDER BY created_at DESC;
```

```sql
/* @name CreateUser @exec */
INSERT INTO users (email, name) VALUES (:email, :name);
```

```sql
/* @name GetUsersByIDs @many */
SELECT * FROM users WHERE id IN (:ids...);
```

**Generate:**

```bash
sqlty
# *chef's kiss*
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

    // Your IDE knows what this returns. Your compiler enforces it.
    // Your deploy doesn't page you at 3am. Beautiful.
    users, _ := q.ListActiveUsers(ctx)

    // Spread parameters just work
    users, _ = q.GetUsersByIDs(ctx, []int64{1, 2, 3})

    // Transactions are context-based because we're not barbarians
    tx, _ := conn.Begin(ctx)
    ctx = context.WithValue(ctx, db.CtxDBKey, tx)
    q.CreateUser(ctx, "alice@example.com", "Alice")
    tx.Commit(ctx)
}
```

---

## The Annotation Cheat Sheet

| Annotation                | What It Does                                          |
| ------------------------- | ----------------------------------------------------- |
| `@name FunctionName`      | Names your function (shocking, I know)                |
| `@one`                    | Returns one row or `nil`                              |
| `@many`                   | Returns `[]Rows`                                      |
| `@exec`                   | Returns nothing. For INSERTs you don't care about     |
| `@template cursor`        | Streaming for when you SELECT'd too much              |
| `@param p (...)`          | `WHERE id IN (:ids...)` without counting placeholders |
| `@param p ((a,b)...)`     | Bulk inserts that don't make you hate your life       |
| `@notNullParams a, b`     | "Trust me, these nullable columns won't be null here" |
| `@returnValueName Custom` | Name your return struct something prettier            |
| `@paramStructName Custom` | Name your params struct (for the aesthetically picky) |
| `:param!`                 | Inline not-null: "this one specifically, trust me"    |

---

## What Gets Generated

```
db/
├── get_user_by_email.gen.go     # Your queries, but Go
├── list_active_users.gen.go
├── enums.sqlty.gen.go           # CREATE TYPE ... AS ENUM → Go const
├── composite_types.sqlty.gen.go  # Postgres composite types → Go structs
└── db.sqlty.gen.go              # The glue
```

Commit these. Or don't. I'm not your dad.

---

## Requirements

- Go 1.21+
- PostgreSQL (MySQL support left as an exercise for the reader) (PRs welcome)
  (please)
- pgx/v5

---

## FAQ

**Q: Is this production ready?** A: I run it in production. Draw your own
conclusions.

**Q: How does it compare to sqlc?** A: sqlc is great. SQLty exists because I
wanted something slightly different and had a free weekend. Use what works for
you.

**Q: Why PostgreSQL only?** A: Because Postgres is the only database. Everything
else is just temporary storage with delusions of grandeur.

**Q: Will you add [feature]?** A: Open an issue. I'm susceptible to peer
pressure.

---

## License

MIT — Do whatever you want. I'm not a cop.

---

<p align="center">
  <i>Built with caffeine and mass-assigned foot-guns</i>
</p>
