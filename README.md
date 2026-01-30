# SQLty

SQLty is a SQL code generator for Go that transforms annotated SQL query files into type-safe Go functions. It connects to a PostgreSQL database to resolve parameter and return types, then generates Go code with compile-time type safety.

## Features

- **Type-safe SQL**: Generate Go functions with proper types resolved from your PostgreSQL schema
- **Multiple execution modes**: `@one` (single row), `@many` (multiple rows), `@exec` (no return)
- **Cursor support**: Stream large result sets with lazy evaluation
- **Composite types**: Full support for PostgreSQL composite/row types
- **Enum types**: Automatic Go enum generation from PostgreSQL enums
- **Spread parameters**: Expand arrays into query parameters (`:ids...`)
- **Struct spread**: Bulk insert support with struct arrays

## Requirements

- Go 1.21+
- PostgreSQL database
- `goimports` (automatically run on generated code)

## Installation

```bash
go install github.com/ioj/sqlty@latest
```

## Usage

1. Create a `sqlty.yaml` configuration file:

```yaml
dburl: postgres://user:pass@localhost:5432/mydb
paths:
  source: ./queries
  output: ./db
package: db
```

2. Write annotated SQL files in your source directory:

```sql
/* @name GetUser
   @one
*/
SELECT id, email, name FROM users WHERE id = :userId;

/* @name ListUsers
   @many
*/
SELECT id, email, name FROM users ORDER BY id;

/* @name CreateUser
   @exec
*/
INSERT INTO users (email, name) VALUES (:email, :name);
```

3. Run SQLty:

```bash
sqlty
```

4. Use the generated code:

```go
import "yourproject/db"

func main() {
    conn, _ := pgx.Connect(ctx, connString)
    queries := db.New(conn)

    user, err := queries.GetUser(ctx, 123)
    users, err := queries.ListUsers(ctx)
}
```

## SQL Annotations

| Annotation | Description |
|------------|-------------|
| `@name queryName` | Name of the generated Go function |
| `@one` | Returns a single row (uses `CollectOneRow`) |
| `@many` | Returns multiple rows (uses `CollectRows`) |
| `@exec` | No return value (INSERT/UPDATE/DELETE) |
| `@param name -> type` | Override parameter type |
| `@notNull paramName` | Mark parameter as non-nullable |
| `@template templateName` | Use custom template (e.g., `cursor`) |
| `@return StructName` | Custom name for return struct |

## Generated Output

SQLty generates the following files:

- `{query_name}.gen.go` - Query functions
- `enums.sqlty.gen.go` - PostgreSQL enum types
- `composite_types.sqlty.gen.go` - Composite/row types
- `db.sqlty.gen.go` - DB struct and utilities

---

## pgx v5 Migration (v2.0)

SQLty has been upgraded from pgx/v4 to pgx/v5, bringing significant improvements to the generated code and type system.

### Breaking Changes

- **Minimum Go version**: Now requires Go 1.21+
- **Type mappings changed**:
  - JSON/JSONB: `pgtype.JSON`/`pgtype.JSONB` → `[]byte`
  - Network types: `pgtype.CIDR`/`pgtype.Inet` → `netip.Prefix`/`netip.Addr`
  - Bit type: `pgtype.Bit` → `pgtype.Bits`
- **Removed DecodeBinary**: Composite types now use pgx v5's native struct scanning

### New Features

#### CollectRows Helper

`@many` queries now use pgx v5's `CollectRows` for cleaner, more efficient code:

```go
// Before (v4)
rows, err := db.tx.Query(ctx, stmt, args...)
defer rows.Close()
var results []*User
for rows.Next() {
    r := &User{}
    if err := rows.Scan(&r.Id, &r.Email); err != nil {
        return nil, err
    }
    results = append(results, r)
}
return results, rows.Err()

// After (v5)
rows, err := db.tx.Query(ctx, stmt, args...)
if err != nil {
    return nil, err
}
return pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[User])
```

#### CollectOneRow Helper

`@one` queries use `CollectOneRow` for simpler single-row retrieval:

```go
// Before (v4)
result := &User{}
err := db.tx.QueryRow(ctx, stmt, args...).Scan(&result.Id, &result.Email)

// After (v5)
rows, err := db.tx.Query(ctx, stmt, args...)
if err != nil {
    return nil, err
}
return pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByPos[User])
```

#### Simplified Connection

Internal type resolution now uses a single `pgx.Conn` instead of separate pgconn and pgx connections.

#### Modern Error Handling

Error code extraction uses the `errors.As` pattern:

```go
func PGErrCode(err error) string {
    var pgerr *pgconn.PgError
    if errors.As(err, &pgerr) {
        return pgerr.Code
    }
    return ""
}
```

### Type Mapping Changes

| PostgreSQL Type | v4 Go Type | v5 Go Type |
|-----------------|------------|------------|
| `json` | `pgtype.JSON` | `[]byte` |
| `jsonb` | `pgtype.JSONB` | `[]byte` |
| `cidr` | `pgtype.CIDR` / `net.IPNet` | `netip.Prefix` |
| `inet` | `pgtype.Inet` | `netip.Addr` |
| `bit` | `pgtype.Bit` | `pgtype.Bits` |
| `timestamptz` (nullable) | `pgtype.Timestamp` | `pgtype.Timestamptz` |
| `oid`, `cid`, `xid` | `pgtype.OID`, etc. | `uint32` |

### Cursor Queries

Cursor queries (`@template cursor`) continue to use manual iteration since they require lazy evaluation:

```go
cursor, err := queries.ListUsersCursor(ctx)
if err != nil {
    return err
}
defer cursor.Close()

for cursor.Next() {
    user, err := cursor.Scan()
    if err != nil {
        return err
    }
    // Process user...
}
return cursor.Err()
```

### Migration Guide

1. **Update your Go version** to 1.21 or later

2. **Update dependencies**:
   ```bash
   go get github.com/ioj/sqlty@latest
   go mod tidy
   ```

3. **Regenerate code**:
   ```bash
   sqlty
   ```

4. **Update imports** in your code if you were using pgtype types directly:
   ```go
   // Before
   import "github.com/jackc/pgtype"

   // After
   import "github.com/jackc/pgx/v5/pgtype"
   ```

5. **Update type usage** for changed mappings:
   ```go
   // JSON handling - now []byte
   var data []byte
   err := json.Unmarshal(row.JsonColumn, &myStruct)

   // Network types - now netip
   import "net/netip"
   var addr netip.Addr
   var prefix netip.Prefix
   ```

---

## Configuration

Create `sqlty.yaml` in your project root:

```yaml
# Database connection string (required)
dburl: postgres://user:pass@localhost:5432/mydb

# File paths
paths:
  source: ./queries    # SQL files location
  output: ./db         # Generated Go files location
  cache: .sqlty        # Cache directory for incremental builds
  templates: ""        # Custom templates directory (optional)

# Package name for generated code
package: db

# Custom type mappings (optional)
types:
  - name: pg_catalog.money
    notNull: true
    to:
      name: "decimal.Decimal"
      zeroValue: "decimal.Decimal{}"
      nullable: false
```

## License

MIT
