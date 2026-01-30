# Generator Package

The `generator` package transforms query definitions and type information into
Go source code. It operates as the final stage in the SQLty pipeline, taking
processed query and type data and generating type-safe Go functions using Go's
`text/template` system.

## Overview

```
stmt.Query/Enums/CompositeTypes/DB → Generator → .gen.go files
```

The generator:

- Loads embedded templates (with optional custom overrides)
- Checks a cache to skip unchanged files
- Executes templates with query/type data
- Writes `.gen.go` files to the output directory

## Files

| File           | Purpose                                                       |
| -------------- | ------------------------------------------------------------- |
| `generator.go` | Core generator logic, template functions, and file generation |
| `cache.go`     | Incremental build caching using MD5 checksums                 |
| `templates/`   | Embedded Go templates for code generation                     |

## Usage

```go
// Create generator with optional custom templates and cache directory
gen, err := generator.New(templateDir, cacheDir)
if err != nil {
    return err
}
defer gen.Close()

// Generate DB utilities (db.sqlty.gen.go)
gen.DB(outputPath, &stmt.DB{PackageName: "queries"})

// Generate enums (enums.sqlty.gen.go)
gen.Enums(outputPath, enums)

// Generate composite types (composite_types.sqlty.gen.go)
gen.CompositeTypes(outputPath, compositeTypes)

// Generate query files ({name}.gen.go)
gen.Query("array", outputFile, query)  // Uses query-array.go.tpl
gen.Query("cursor", outputFile, query) // Uses query-cursor.go.tpl
```

## Templates

Templates are embedded in the binary via `//go:embed` and can be overridden by
placing custom templates in the configured template directory.

| Template                 | Output File                    | Purpose                                               |
| ------------------------ | ------------------------------ | ----------------------------------------------------- |
| `db.sqlty.go.tpl`        | `db.sqlty.gen.go`              | DB struct, context utilities, error types, middleware |
| `enums.go.tpl`           | `enums.sqlty.gen.go`           | PostgreSQL enum types as Go constants with index maps |
| `composite_types.go.tpl` | `composite_types.sqlty.gen.go` | PostgreSQL composite types with binary decoding       |
| `structs.go.tpl`         | (included by others)           | Struct definitions and DecodeBinary methods           |
| `query-array.go.tpl`     | `{name}.gen.go`                | Query functions for `@one`, `@many`, `@exec` modes    |
| `query-cursor.go.tpl`    | `{name}.gen.go`                | Cursor-based query functions for streaming results    |

### Template Functions

Custom functions available in templates:

| Function                    | Description                                                   |
| --------------------------- | ------------------------------------------------------------- |
| `firstParamTypeName`        | Returns the type name of the first return parameter           |
| `firstParamNilReturnValue`  | Returns nil or zero value for nullable first param            |
| `firstParamZeroReturnValue` | Returns zero value for first param                            |
| `needsPrintf`               | True if query has spread parameters (requires formatting)     |
| `hasParams`                 | True if query has any parameters                              |
| `valueToIdent`              | Converts SQL values to valid Go identifiers                   |
| `returnsStruct`             | True if query returns a composite or inline struct            |
| `returnsInlineStruct`       | True if query returns multiple columns (not a composite type) |
| `lowerFirstLetter`          | Lowercases first letter (for JSON tags)                       |

## Caching

The generator uses MD5 checksums to implement incremental builds. Before
generating each file, it:

1. Serializes the input parameters to JSON
2. Computes an MD5 checksum
3. Compares against the cached checksum
4. Skips generation if unchanged

Cache is stored in `cache.yaml` within the configured cache directory:

```yaml
items:
  /path/to/output/query.gen.go: "abc123..."
  /path/to/output/db.sqlty.gen.go: "def456..."
```

Benefits:

- Avoids unnecessary file writes
- Preserves file timestamps when unchanged
- Enables fast incremental builds

## Generated Output

### DB Utilities (`db.sqlty.gen.go`)

```go
// Context key for transaction storage
const CtxDBKey = CtxKey(23421)

// Error types
var ErrNilArgs, ErrEmptySpread, ErrEmptyStructSpreadList

// DB struct wrapping a Txer interface
type DB struct { tx Txer }
func New(db Txer) *DB

// HTTP middleware for transaction-per-request
func Middleware(db Beginner) func(http.Handler) http.Handler

// Get transaction from context
func CtxTx(ctx context.Context) (*DB, func(context.Context) error)

// PostgreSQL error code extraction
func PGErrCode(err error) string
```

### Enums (`enums.sqlty.gen.go`)

```go
type Status string

const (
    Status_Active  Status = "active"
    Status_Pending Status = "pending"
)

var StatusIdxMap = map[Status]int{...}
var StatusByIdx = []Status{...}
```

### Query Functions

Generated query methods on the `DB` type handle:

- Parameter binding (scalar, spread, struct spread)
- Execution mode (`@exec`, `@one`, `@many`)
- Return type scanning (single value, struct, slice)
- Spread parameter validation

## Custom Templates

To customize generated code, create `.go.tpl` files in your template directory.
Custom templates override embedded defaults of the same name.

```bash
# In sqlty.yaml
paths:
  templates: ./custom_templates
```

Then create files like `custom_templates/query-array.go.tpl` to override the
default query generation.
