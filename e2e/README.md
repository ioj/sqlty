# End-to-End Tests

Tests that verify sqlty-generated Go functions work correctly against a real
PostgreSQL database.

## Prerequisites

A running PostgreSQL instance with a test database:

```bash
createdb sqlty_test
```

## Running

```bash
TEST_DATABASE_URL="postgres://localhost/sqlty_test" go test -v ./...
```

Tables and the `e2e_user_status` enum are created automatically by the test
setup and dropped on cleanup.

## Structure

- `queries/` — annotated SQL files (sqlty input)
- `generated/` — checked-in generated Go code (sqlty output)
- `e2e_test.go` — tests

## Regenerating

If you change SQL files, regenerate with:

```bash
SQLTY_DBURL="$TEST_DATABASE_URL" sqlty -config sqlty.yaml
```

The tables must exist in the database before running sqlty (run the tests once
first, or create them manually).

## Coverage

| Area            | Tests                                      |
|-----------------|--------------------------------------------|
| CRUD            | CreateAndGetUser, ListUsers, Update, Delete |
| Nullable fields | GetUserNullableEmail                        |
| Spread params   | FindUsersByIds (`IN` clause)                |
| Not-null params | GetUserEmail (`:email!`)                    |
| Enums           | GetUserStatus                               |
| Arrays          | ArrayColumns (int[], text[], uuid[], bool[])|
| Array params    | FindByTags (`&&` overlap)                   |
