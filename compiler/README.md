# SQLty Compiler

The compiler package parses SQL files with special annotations and extracts
metadata needed for Go code generation. It uses a hand-written lexer and
recursive descent parser.

## Architecture Overview

```
SQL File Input
    │
    ▼
┌─────────────────┐
│     Lexer       │  Tokenizes input with mode switching
│   (lexer.go)    │  DEFAULT mode ↔ COMMENT mode
└────────┬────────┘
         │ Token stream
         ▼
┌─────────────────┐
│     Parser      │  Recursive descent parser
│   (parser.go)   │  Builds Query struct from tokens
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│     Query       │  Complete metadata for code generation
│   (query.go)    │  Parameters, exec mode, statement text
└─────────────────┘
```

## File Structure

| File          | Description                                         |
| ------------- | --------------------------------------------------- |
| `compiler.go` | Entry points: `CompileFile()` and `CompileString()` |
| `lexer.go`    | Tokenizer with dual-mode lexing                     |
| `parser.go`   | Recursive descent parser                            |
| `query.go`    | Data structures (`Query`, `Param`, `token`)         |
| `errors.go`   | Error types and reporting                           |

## Lexer (`lexer.go`)

The lexer tokenizes SQL input using two modes:

### DEFAULT Mode (SQL body)

Used when parsing the SQL statement itself.

| Token              | Pattern                  | Description                           |
| ------------------ | ------------------------ | ------------------------------------- |
| `TokenOpenComment` | `/*`                     | Switches to COMMENT mode              |
| `TokenLineComment` | `-- ... \n`              | Line comment (docstring)              |
| `TokenParamMark`   | `:`                      | Parameter marker                      |
| `TokenIdentifier`  | `[a-zA-Z_][a-zA-Z0-9_]*` | Identifier                            |
| `TokenWord`        | any sequence             | SQL keywords, operators               |
| `TokenString`      | `'...'`                  | String literal (handles `''` escapes) |
| `TokenSemicolon`   | `;`                      | Statement terminator                  |
| `TokenPercent`     | `%`                      | Percent sign (for sprintf escaping)   |

### COMMENT Mode (annotation block)

Used inside `/* ... */` to parse annotation tags.

| Token                    | Pattern            | Description               |
| ------------------------ | ------------------ | ------------------------- |
| `TokenCloseComment`      | `*/`               | Returns to DEFAULT mode   |
| `TokenAtName`            | `@name`            | Query name tag            |
| `TokenAtParam`           | `@param`           | Parameter declaration tag |
| `TokenAtParamStructName` | `@paramStructName` | Param struct wrapper name |
| `TokenAtOne`             | `@one`             | Return single row         |
| `TokenAtMany`            | `@many`            | Return multiple rows      |
| `TokenAtExec`            | `@exec`            | Execute without return    |
| `TokenAtNotNullParams`   | `@notNullParams`   | Not-null constraints      |
| `TokenAtReturnValueName` | `@returnValueName` | Return struct name        |
| `TokenAtTemplate`        | `@template`        | Template selection        |
| `TokenOpenParen`         | `(`                | Open parenthesis          |
| `TokenCloseParen`        | `)`                | Close parenthesis         |
| `TokenComma`             | `,`                | Comma separator           |
| `TokenDot`               | `.`                | Dot (for `param.field`)   |
| `TokenSpread`            | `...`              | Spread operator           |
| `TokenArrow`             | `->`               | Arrow (param transform)   |

### Mode Switching

```
┌─────────────┐    /*     ┌─────────────┐
│   DEFAULT   │ ────────► │   COMMENT   │
│    MODE     │ ◄──────── │    MODE     │
└─────────────┘    */     └─────────────┘
```

## Parser (`parser.go`)

The parser is a recursive descent parser that builds a `Query` struct.

### Parsing Flow

1. **Annotation Block** (`parseAnnotationBlock`)
   - Expects `/*` at start
   - Parses tags until `*/`

2. **Line Comments** (`parseLineComments`)
   - Collects `--` comments as docstrings
   - Only comments before the SQL statement are captured

3. **Statement Body** (`parseStatementBody`)
   - Tracks parameter usages (`:paramName`)
   - Tracks percent signs for sprintf escaping
   - Captures full SQL text

4. **Validation**
   - `checkUnusedParameters()` - Error if `@param` declared but not used
   - `populateNotNullParams()` - Links `@notNullParams` to params
   - `verifyExecMode()` - Error if `@one/@many/@exec` missing

### Annotation Tags

#### `@name`

Sets the query function name.

```sql
/* @name GetUserByID */
```

#### `@param`

Declares a parameter with optional type transform.

**Scalar** (default): Single value parameter

```sql
/* @param id */
SELECT * FROM users WHERE id = :id;
```

**Spread**: Expands to multiple values for `IN` clauses

```sql
/* @param ids -> (...) */
SELECT * FROM users WHERE id IN :ids;
```

**Struct Spread**: Expands structs for bulk inserts

```sql
/* @param users -> ((name, email)...) */
INSERT INTO users (name, email) VALUES :users;
```

#### `@one` / `@many` / `@exec`

Specifies the execution mode (required).

- `@one` - Query returns a single row
- `@many` - Query returns multiple rows
- `@exec` - Execute without return value

#### `@notNullParams`

Marks parameters or struct fields as non-nullable.

```sql
/* @notNullParams (id, user.name) */
```

#### `@returnValueName`

Overrides the default return struct name.

```sql
/* @returnValueName UserRow */
```

#### `@paramStructName`

Groups all parameters into a struct.

```sql
/* @paramStructName GetUserParams */
```

#### `@template`

Selects a custom code generation template.

```sql
/* @template custom */
```

## Query Structure (`query.go`)

### Core Types

```go
type Query struct {
    Filename        string           // Source file
    Comments        []string         // Line comment docstrings
    name            *token           // @name value
    execMode        *token           // @one/@many/@exec
    template        *token           // @template value
    params          map[string]*Param
    paramStructName *token           // @paramStructName value
    returnValueName *token           // @returnValueName value
    notNullParams   []*token         // @notNullParams list
    percents        []*token         // % positions in SQL
    statement       *token           // The SQL statement
}

type Param struct {
    definition *token              // Where declared
    uses       []*token            // All :name usages in SQL
    keys       map[string]*StructKey  // For struct spreads
    Type       ParamType           // Scalar/Spread/StructSpread
    Idx        int                 // Index within type group
    NotNull    bool
}

type ParamType string
const (
    Scalar       ParamType = "scalar"
    Spread       ParamType = "spread"
    StructSpread ParamType = "structspread"
)
```

### Key Methods

- `Query.Statement()` - Returns SQL with `:params` replaced by `$N` placeholders
- `Query.PreparedQuery()` - Returns SQL ready for PostgreSQL `Prepare()`
- `Query.StmtQuery()` - Converts to `stmt.Query` for code generation

## Error Handling (`errors.go`)

Errors are accumulated during parsing and returned as `ErrCompilationFailed`.

| Error Type   | Description                |
| ------------ | -------------------------- |
| `syntax`     | Lexer/parser syntax error  |
| `annotation` | Invalid annotation usage   |
| `empty`      | Empty or missing statement |

Error messages include file, line, and column for precise reporting:

```
[error][annotation] query.sql:5:2: parameter `id` is declared, but not used in the query
```

## Usage Example

```go
// Parse a file
query, err := compiler.CompileFile("queries/get_user.sql")
if err != nil {
    log.Fatal(err)
}

// Parse a string
query, err := compiler.CompileString("query", `
/* @name GetUser @one */
SELECT * FROM users WHERE id = :id;
`)
```

## Example SQL File

```sql
/*
  @name GetUsersByStatus
  @param status
  @param limit
  @notNullParams (status)
  @many
  @returnValueName User
*/
-- GetUsersByStatus retrieves users filtered by status.
-- Returns up to `limit` results.
SELECT id, name, email, status
FROM users
WHERE status = :status
ORDER BY created_at DESC
LIMIT :limit;
```

This produces:

- Function name: `GetUsersByStatus`
- Parameters: `status` (not-null), `limit`
- Return type: `[]User`
- Docstring from line comments
