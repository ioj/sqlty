# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with
code in this repository.

## Project Overview

SQLty is a SQL code generator for Go that transforms annotated SQL query files
into type-safe Go functions. It connects to a PostgreSQL database to resolve
parameter and return types, then generates Go code with compile-time type
safety.

## Build and Development Commands

```bash
# Build
go build ./...

# Run all tests
go test ./...

# Run a single test
go test -run TestSimpleSelect ./compiler
```

Note: `goimports -w .` is automatically run on the output directory after code
generation.

## Architecture

SQLty follows a pipeline architecture with these main stages:

### 1. Config (`config/`)

Loads configuration from `sqlty.yaml`. Key settings: database URL, source/output
paths, package name, type mappings.

### 2. Compiler (`compiler/`)

Parses SQL files using a hand-written lexer and recursive descent parser to
extract query metadata:

- Query name and execution mode (`@one`, `@many`, `@exec`)
- Parameters (with types like scalar, spread, struct spread)
- Template selection and return struct names
- Not-null constraints

Key files: `lexer.go` (tokenizer), `parser.go` (recursive descent parser),
`query.go` (AST types).

### 3. DB Resolver (`db/`)

Connects to PostgreSQL and resolves types:

- Maps SQL parameter/return OIDs to Go types
- Discovers composite types and enums from system catalogs (`pg_type`,
  `pg_enum`, `pg_attribute`)
- Type mappings defined in embedded `types.yaml`

### 4. Statement Model (`stmt/`)

Data structures representing the generated code units: Query, Param, Type,
Struct, Enum, ExecMode.

### 5. Generator (`generator/`)

Converts query and type info into Go code using embedded templates
(`generator/templates/`). Supports custom template overrides and incremental
builds via caching.

## Execution Flow

```
SQL File → Compiler (parse annotations) → DB Resolver (resolve types) → Statement Model → Generator → Go Code
```

## SQL Annotation Syntax

```sql
/* @name queryName
   @param paramName (...)
   @notNull paramName
   @template templateName
   @exec @one/@many/@exec
   @return returnStructName
*/
SELECT * FROM table WHERE id = :paramName;
```

## Generated Output

- `{query_name}.gen.go` - Query functions
- `enums.sqlty.gen.go` - PostgreSQL enum types
- `composite_types.sqlty.gen.go` - Composite/row types
- `db.sqlty.gen.go` - DB struct and utilities

## Key Patterns

- **Context transactions**: Uses `CtxDBKey` constant to attach transactions to
  contexts
- **Error handling**: Custom error types in `compiler/errors.go` for validation
- **Recursive descent parsing**: Hand-written lexer/parser in `compiler/`
- **Embedding**: Default templates and type mappings bundled in binary via
  `//go:embed`
