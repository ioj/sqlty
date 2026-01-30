# stmt Package

The `stmt` package defines the data model that bridges SQL parsing and Go code
generation. It provides intermediate data structures that capture all the
information needed to generate type-safe Go code from SQL queries.

## Overview

This package contains no business logic - it defines pure data structures that
flow through SQLty's pipeline:

```
Compiler (SQL parsing) → stmt types → Generator (Go code output)
```

## Types

### Query

The primary type representing a parsed SQL query ready for code generation.

```go
type Query struct {
    PackageName string    // Target Go package name
    Name        string    // Generated function name
    Statement   string    // The SQL query text
    ExecMode    ExecMode  // How to execute: exec, one, or many
    Comments    []string  // Doc comments for the generated function
    Params      Params    // Input parameters
    Returns     Struct    // Return type structure
}
```

### ExecMode

Determines how the generated function executes the query:

- `exec` - Execute without returning rows (INSERT, UPDATE, DELETE)
- `one` - Return a single row
- `many` - Return multiple rows

### Params

Organizes query parameters by their binding style:

```go
type Params struct {
    Name         string    // If set, params are wrapped in a struct
    Scalar       []Param   // Single value parameters (:id)
    Spread       []Param   // Expanded into multiple placeholders (:ids...)
    StructSpread []Struct  // Struct fields expanded as parameters
}
```

### Param and Type

A parameter with its resolved Go type:

```go
type Param struct {
    Name string
    Type Type
}

type Type struct {
    Name      string  // Go type name (e.g., "int64", "string")
    ZeroValue string  // Zero value for the type (e.g., "0", `""`)
    Nullable  bool    // Whether the type can be null
}
```

### Struct

Represents a Go struct for return types or composite parameters:

```go
type Struct struct {
    Name            string
    Params          []Param
    IsCompositeType bool  // True if this maps to a PostgreSQL composite type
}
```

### Enum

Represents a PostgreSQL enum type:

```go
type Enum struct {
    Name   string    // Go type name
    Values []string  // Enum values
}
```

### Container Types

Types that group multiple items for generating dedicated files:

- `Enums` - Collection of enums for `enums.sqlty.gen.go`
- `CompositeTypes` - Collection of composite types for
  `composite_types.sqlty.gen.go`
- `DB` - Database struct metadata for `db.sqlty.gen.go`

## Usage

These types are populated by the compiler and db resolver, then passed to the
generator's templates. Each type includes the `PackageName` field needed for the
generated Go package declaration.
