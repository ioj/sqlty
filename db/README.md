# db Package

The `db` package is SQLty's PostgreSQL type resolver. It connects to a
PostgreSQL database to discover type information and maps PostgreSQL types to
their Go equivalents.

## Purpose

When SQLty compiles SQL queries, it needs to determine:

1. What Go types should parameters have?
2. What Go types should return values have?

This package answers those questions by:

- Preparing queries against a real PostgreSQL database
- Inspecting the parameter and return type OIDs from PostgreSQL
- Mapping those OIDs to appropriate Go types

## Core Components

### Resolver

The main entry point. Create one with `NewResolver()`, use it to resolve types,
then call `Close()`.

```go
resolver, err := db.NewResolver(ctx, connectionString, customTypeMappings)
defer resolver.Close()

// Resolve parameter and return types for a query
params, returns, err := resolver.ResolveTypes(ctx, "SELECT * FROM users WHERE id = $1", []bool{true})

// Get all PostgreSQL enums (for code generation)
enums := resolver.Enums()

// Get all composite types used by queries
composites, err := resolver.CompositeTypes(ctx)
```

### pgTypes

Internal type registry that:

- Loads all types from `pg_type` system catalog
- Loads enum labels from `pg_enum`
- Loads composite type fields from `pg_attribute`
- Maintains a translation table from PostgreSQL types to Go types

### types.yaml

Embedded YAML file containing default PostgreSQL-to-Go type mappings. Each
mapping specifies:

- `name`: Full PostgreSQL type name (e.g., `pg_catalog.int4`)
- `notNull`: Whether this mapping applies to NOT NULL columns
- `to`: The Go type details (name, zero value, nullable flag)

## Type Resolution Flow

```
SQL Query
    │
    ▼
PostgreSQL PREPARE
    │
    ▼
Extract Parameter OIDs + Return Field OIDs
    │
    ▼
Look up OIDs in pg_type cache
    │
    ▼
Match against type translations
    │
    ▼
Return Go types (stmt.Type)
```

## PostgreSQL System Catalogs Used

| Catalog        | Purpose                                                 |
| -------------- | ------------------------------------------------------- |
| `pg_type`      | Type definitions (OID, name, category, base type, etc.) |
| `pg_namespace` | Schema/namespace names for fully-qualified type names   |
| `pg_enum`      | Enum label values in sort order                         |
| `pg_attribute` | Column definitions for composite types                  |

## Type Categories Handled

| PostgreSQL Type                   | Go Mapping Strategy                         |
| --------------------------------- | ------------------------------------------- |
| **Base types** (int4, text, etc.) | Direct lookup in types.yaml                 |
| **Enums**                         | Generated Go string type with constants     |
| **Composite types**               | Generated Go struct with fields             |
| **Domains**                       | Resolved to underlying base type            |
| **Arrays**                        | Element type lookup (prefix `_` in pg_type) |

## Nullability

The resolver tracks nullability from two sources:

1. **Domain constraints**: `typnotnull` in `pg_type` for domain types
2. **Column constraints**: `attnotnull` in `pg_attribute` for table columns

When a type is NOT NULL:

- Primitive types use their base form (`int`, `string`, `bool`)
- When nullable, pointer types are used (`*int`, `*string`, `*bool`)

## Extending Type Mappings

Custom type mappings can be provided via the `types` parameter in
`NewResolver()`. These follow the same structure as `types.yaml`:

```go
customTypes := []db.PGTypeTranslation{
    {
        Fullname: "public.my_custom_type",
        NotNull:  true,
        To: stmt.Type{
            Name:      "MyCustomType",
            ZeroValue: "MyCustomType{}",
            Nullable:  false,
        },
    },
}
```

Custom mappings are merged with the defaults, allowing you to override built-in
types or add support for extensions.

## Key Functions

| Function                                      | Description                                |
| --------------------------------------------- | ------------------------------------------ |
| `NewResolver(ctx, connString, types)`         | Create resolver with database connection   |
| `Resolver.ResolveTypes(ctx, query, notnulls)` | Get parameter and return types for a query |
| `Resolver.Enums()`                            | Get all PostgreSQL enum definitions        |
| `Resolver.CompositeTypes(ctx)`                | Get composite types referenced by queries  |
| `Resolver.Close()`                            | Close database connections                 |

## Error Handling

The resolver returns specific errors for edge cases:

- `errVoid`: Query returns `pg_catalog.void` (no return value)
- `errComposite`: Return type is a composite (triggers struct generation)
- Unknown OID: Returns error with OID value for debugging
