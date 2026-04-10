# E2E Tests for SQLty

## Goal

End-to-end tests that verify generated Go functions return correct values when executed against a real PostgreSQL database. Tests should be minimal, human-readable, and cover the most common paths.

## Structure

```
e2e/
  queries/           # Annotated SQL files (sqlty input)
  generated/         # Checked-in generated Go code (sqlty output)
  e2e_test.go        # Tests that run generated functions against real DB
  go.mod             # Separate module
```

Generated code is checked in. Tests only exercise the generated functions — no pipeline invocation at test time. Regeneration is a manual/CI step.

## Database Schema

```sql
CREATE TYPE user_status AS ENUM ('active', 'inactive', 'banned');

CREATE TABLE e2e_users (
  id         uuid PRIMARY KEY,
  name       text NOT NULL,
  email      text,              -- nullable
  active     bool NOT NULL,
  score      float8 NOT NULL,
  created_at timestamptz NOT NULL,
  status     user_status NOT NULL
);

CREATE TABLE e2e_arrays (
  id        uuid PRIMARY KEY,
  int_tags  int[] NOT NULL,
  str_tags  text[] NOT NULL,
  uuid_refs uuid[] NOT NULL,
  flags     bool[] NOT NULL
);
```

Tables are created in `TestMain` and dropped on teardown.

## SQL Files and Coverage

| File | Mode | Params | Types Covered |
|---|---|---|---|
| `get_user.sql` | `@one` | `:id` (uuid) | uuid, text, *text (nullable), bool, float8, timestamptz, enum |
| `list_users.sql` | `@many` | none | multi-row struct returns |
| `create_user.sql` | `@exec` | all columns | INSERT, command tag |
| `update_user.sql` | `@exec` | `:id`, `:name`, `:score` | UPDATE multiple params |
| `delete_user.sql` | `@exec` | `:id` | DELETE |
| `find_users_by_ids.sql` | `@many` | `@param ids (...)` | spread params, IN clause |
| `get_user_email.sql` | `@one` | `:id`, `:email!` | not-null override |
| `get_arrays.sql` | `@one` | `:id` | array returns: int[], text[], uuid[], bool[] |
| `find_by_tags.sql` | `@many` | `:tags` (text[]) | array param with ANY() |
| `get_user_status.sql` | `@one` | `:id` | enum return type |

## Test Flow

1. `TestMain`: connect to DB, create enum + tables, seed data, run tests, drop tables
2. Helper: `newDB(t)` returns `generated.DB` wrapping a fresh connection
3. Each `Test*` function calls one generated function and asserts the result

## Dependencies

- `github.com/jackc/pgx/v5`
- `github.com/stretchr/testify`
- `TEST_DATABASE_URL` environment variable
