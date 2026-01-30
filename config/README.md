# config Package

The `config` package loads SQLty's configuration from a YAML file and
environment variables.

## Purpose

SQLty needs configuration to know:

- How to connect to the PostgreSQL database
- Where to find SQL source files
- Where to write generated Go code
- What package name to use
- Any custom type mappings

This package handles loading and validating that configuration.

## Configuration File

SQLty looks for `sqlty.yaml` in the current working directory.

```yaml
# Required: PostgreSQL connection string
dbUrl: postgres://user:pass@localhost:5432/mydb

# Optional: Path settings
paths:
  source: ./queries # Where to find .sql files (default: ".")
  output: ./generated # Where to write .go files (default: ".")
  cache: .sqlty # Cache directory for incremental builds (default: ".sqlty")
  templates: ./tpl # Custom template directory (default: "" = use built-in)

# Optional: Go package name for generated code
# Default: derived from output directory name
packageName: mypackage

# Optional: Custom PostgreSQL-to-Go type mappings
types:
  - name: public.my_type
    notNull: true
    to:
      name: MyType
      zeroValue: "MyType{}"
      nullable: false
```

## Configuration Structure

### Config

| Field          | Type                     | Required | Description                        |
| -------------- | ------------------------ | -------- | ---------------------------------- |
| `DBURL`        | `string`                 | Yes      | PostgreSQL connection string       |
| `Paths`        | `Paths`                  | No       | Directory paths (has defaults)     |
| `PackageName`  | `string`                 | No       | Go package name (derived if empty) |
| `Types`        | `[]db.PGTypeTranslation` | No       | Custom type mappings               |
| `DefaultTypes` | `string`                 | No       | Path to default types file         |

### Paths

| Field       | Default    | Description                            |
| ----------- | ---------- | -------------------------------------- |
| `Source`    | `"."`      | Directory containing SQL files         |
| `Output`    | `"."`      | Directory for generated Go files       |
| `Cache`     | `".sqlty"` | Cache directory for incremental builds |
| `Templates` | `""`       | Custom templates directory (optional)  |

## Environment Variables

All configuration values can be overridden via environment variables with the
`SQLTY_` prefix:

| Environment Variable | Config Field   |
| -------------------- | -------------- |
| `SQLTY_DBURL`        | `dbUrl`        |
| `SQLTY_PATHS_SOURCE` | `paths.source` |
| `SQLTY_PATHS_OUTPUT` | `paths.output` |
| `SQLTY_PATHS_CACHE`  | `paths.cache`  |
| `SQLTY_PACKAGENAME`  | `packageName`  |

Environment variables take precedence over file values.

## Usage

```go
cfg, err := config.Load()
if err != nil {
    log.Fatal(err)
}

// Access configuration
fmt.Println(cfg.DBURL)
fmt.Println(cfg.Paths.Source)
fmt.Println(cfg.PackageName)
```

## Package Name Inference

If `packageName` is not specified, it's derived from the output directory:

```
paths.output: ./internal/db/queries
packageName:  queries  (lowercase basename)
```

## Validation

`Load()` returns an error if:

- `sqlty.yaml` is not found in the current directory
- The YAML is malformed
- `dbUrl` is not provided (required field)

## Dependencies

- [spf13/viper](https://github.com/spf13/viper) - Configuration management
