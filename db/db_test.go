package db

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

//go:embed types.yaml
var typesYAML []byte

func getTestConnString(t *testing.T) string {
	t.Helper()
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Fatal("TEST_DATABASE_URL environment variable must be set to run tests")
	}
	return url
}

func loadDefaultTypes(t *testing.T) []PGTypeTranslation {
	t.Helper()
	var typesFile PGTypeTranslationsFile
	err := yaml.NewDecoder(bytes.NewReader(typesYAML)).Decode(&typesFile)
	require.NoError(t, err, "Failed to decode types.yaml")
	return typesFile.Types
}

func setupTestDB(t *testing.T) *pgx.Conn {
	t.Helper()
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, getTestConnString(t))
	require.NoError(t, err, "Failed to connect to test database")
	t.Cleanup(func() { conn.Close(ctx) })
	return conn
}

func setupResolver(t *testing.T) *Resolver {
	t.Helper()
	ctx := context.Background()
	types := loadDefaultTypes(t)
	r, err := NewResolver(ctx, getTestConnString(t), types)
	require.NoError(t, err, "Failed to create resolver")
	t.Cleanup(func() { r.Close() })
	return r
}

// TestNewResolver verifies basic resolver creation and cleanup.
func TestNewResolver(t *testing.T) {
	ctx := context.Background()
	types := loadDefaultTypes(t)
	r, err := NewResolver(ctx, getTestConnString(t), types)
	require.NoError(t, err)
	require.NotNil(t, r)
	require.NoError(t, r.Close())
}

// TestNewResolverInvalidConnection verifies error handling for bad connection strings.
func TestNewResolverInvalidConnection(t *testing.T) {
	ctx := context.Background()
	_, err := NewResolver(ctx, "postgres://invalid:invalid@localhost:9999/nonexistent", nil)
	require.Error(t, err)
}

// TestTypeResolution verifies that PostgreSQL types resolve to correct Go types.
func TestTypeResolution(t *testing.T) {
	r := setupResolver(t)
	ctx := context.Background()

	tests := []struct {
		name     string
		query    string
		notnulls []bool
		wantType string
	}{
		{"int4 not null", "SELECT $1::int4", []bool{true}, "int"},
		{"int4 nullable", "SELECT $1::int4", []bool{false}, "*int"},
		{"int8 not null", "SELECT $1::int8", []bool{true}, "int"},
		{"text not null", "SELECT $1::text", []bool{true}, "string"},
		{"text nullable", "SELECT $1::text", []bool{false}, "*string"},
		{"bool not null", "SELECT $1::bool", []bool{true}, "bool"},
		{"bool nullable", "SELECT $1::bool", []bool{false}, "*bool"},
		{"float4 not null", "SELECT $1::float4", []bool{true}, "float32"},
		{"float8 not null", "SELECT $1::float8", []bool{true}, "float64"},
		{"bytea", "SELECT $1::bytea", []bool{true}, "[]byte"},
		{"uuid", "SELECT $1::uuid", []bool{true}, "pgtype.UUID"},
		{"json", "SELECT $1::json", []bool{true}, "[]byte"},
		{"jsonb", "SELECT $1::jsonb", []bool{true}, "[]byte"},
		{"int2 not null", "SELECT $1::int2", []bool{true}, "int16"},
		{"int2 nullable", "SELECT $1::int2", []bool{false}, "*int16"},
		{"varchar not null", "SELECT $1::varchar", []bool{true}, "string"},
		{"numeric", "SELECT $1::numeric", []bool{true}, "pgtype.Numeric"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params, _, err := r.ResolveTypes(ctx, tt.query, tt.notnulls)
			require.NoError(t, err)
			require.Len(t, params, 1)
			assert.Equal(t, tt.wantType, params[0].Name)
		})
	}
}

// TestReturnTypes verifies that return types are resolved correctly.
// Note: Literal values (not from table columns) are nullable by default.
func TestReturnTypes(t *testing.T) {
	r := setupResolver(t)
	ctx := context.Background()

	tests := []struct {
		name      string
		query     string
		wantCount int
		wantTypes []string
	}{
		// Literal values are nullable since there's no column constraint
		{"single int", "SELECT 1::int4", 1, []string{"*int"}},
		{"multiple columns", "SELECT 1::int4, 'hello'::text", 2, []string{"*int", "*string"}},
		{"three columns", "SELECT true::bool, 1.5::float8, 'x'::varchar", 3, []string{"*bool", "*float64", "*string"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, returns, err := r.ResolveTypes(ctx, tt.query, nil)
			require.NoError(t, err)
			require.NotNil(t, returns)
			require.Len(t, returns.Params, tt.wantCount)
			for i, wantType := range tt.wantTypes {
				assert.Equal(t, wantType, returns.Params[i].Type.Name)
			}
		})
	}
}

// TestVoidReturn verifies that void-returning functions are handled correctly.
func TestVoidReturn(t *testing.T) {
	r := setupResolver(t)
	ctx := context.Background()

	// pg_sleep returns void
	_, returns, err := r.ResolveTypes(ctx, "SELECT pg_sleep(0)", nil)
	require.NoError(t, err)
	assert.Nil(t, returns, "void return should produce nil returns")
}

// TestParamCountMismatch verifies error when notnulls array doesn't match params.
func TestParamCountMismatch(t *testing.T) {
	r := setupResolver(t)
	ctx := context.Background()

	// Query has 2 params, but we only provide 1 notnull
	_, _, err := r.ResolveTypes(ctx, "SELECT $1::int4, $2::text", []bool{true})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid number of parameters")
}

// TestCustomTypeTranslation verifies that custom type mappings override defaults.
// Custom types must come after defaults since later entries overwrite earlier ones.
func TestCustomTypeTranslation(t *testing.T) {
	ctx := context.Background()
	// Start with default types, then add custom override at the end
	defaultTypes := loadDefaultTypes(t)
	customTypes := append(defaultTypes, PGTypeTranslation{
		Fullname: "pg_catalog.int4",
		NotNull:  true,
		To: struct {
			Name      string "yaml:\"name\""
			ZeroValue string "yaml:\"zeroValue\""
			Nullable  bool   "yaml:\"nullable\""
		}{Name: "CustomInt", ZeroValue: "0", Nullable: false},
	})

	r, err := NewResolver(ctx, getTestConnString(t), customTypes)
	require.NoError(t, err)
	defer r.Close()

	params, _, err := r.ResolveTypes(ctx, "SELECT $1::int4", []bool{true})
	require.NoError(t, err)
	require.Len(t, params, 1)
	assert.Equal(t, "CustomInt", params[0].Name)
}

// TestInvalidSQL verifies error handling for malformed SQL.
func TestInvalidSQL(t *testing.T) {
	r := setupResolver(t)
	ctx := context.Background()

	_, _, err := r.ResolveTypes(ctx, "SELECT * FROM nonexistent_table_xyz_123", nil)
	require.Error(t, err)
}

// TestEnumsWithTempEnum tests enum discovery by creating a temporary enum.
func TestEnumsWithTempEnum(t *testing.T) {
	conn := setupTestDB(t)
	ctx := context.Background()

	enumName := "test_status_" + randomSuffix()
	_, err := conn.Exec(ctx, "CREATE TYPE "+enumName+" AS ENUM ('pending', 'active', 'done')")
	require.NoError(t, err)
	t.Cleanup(func() {
		conn.Exec(ctx, "DROP TYPE IF EXISTS "+enumName)
	})

	r := setupResolver(t)
	enums := r.Enums()

	var found bool
	for _, e := range enums {
		if e.Name == snakeToCamel(enumName) {
			found = true
			assert.Equal(t, []string{"pending", "active", "done"}, e.Values)
			break
		}
	}
	assert.True(t, found, "Created enum should be discovered")
}

// TestEnumTypeResolution tests that enum parameters resolve correctly.
func TestEnumTypeResolution(t *testing.T) {
	conn := setupTestDB(t)
	ctx := context.Background()

	enumName := "test_priority_" + randomSuffix()
	_, err := conn.Exec(ctx, "CREATE TYPE "+enumName+" AS ENUM ('low', 'medium', 'high')")
	require.NoError(t, err)
	t.Cleanup(func() {
		conn.Exec(ctx, "DROP TYPE IF EXISTS "+enumName)
	})

	r := setupResolver(t)
	goName := snakeToCamel(enumName)

	// Test not null enum
	params, _, err := r.ResolveTypes(ctx, "SELECT $1::"+enumName, []bool{true})
	require.NoError(t, err)
	require.Len(t, params, 1)
	assert.Equal(t, goName, params[0].Name)

	// Test nullable enum
	params, _, err = r.ResolveTypes(ctx, "SELECT $1::"+enumName, []bool{false})
	require.NoError(t, err)
	require.Len(t, params, 1)
	assert.Equal(t, "*"+goName, params[0].Name)
}

// TestCompositeType tests composite type resolution.
func TestCompositeType(t *testing.T) {
	conn := setupTestDB(t)
	ctx := context.Background()

	typeName := "test_person_" + randomSuffix()
	_, err := conn.Exec(ctx, "CREATE TYPE "+typeName+" AS (name text, age int4)")
	require.NoError(t, err)
	t.Cleanup(func() {
		conn.Exec(ctx, "DROP TYPE IF EXISTS "+typeName)
	})

	r := setupResolver(t)

	// Query returning composite type - using ROW constructor
	_, returns, err := r.ResolveTypes(ctx, "SELECT ROW('John', 30)::"+typeName, nil)
	require.NoError(t, err)
	require.NotNil(t, returns)
	assert.True(t, returns.IsCompositeType)
	require.Len(t, returns.Params, 2)
	assert.Equal(t, "Name", returns.Params[0].Name)
	assert.Equal(t, "Age", returns.Params[1].Name)
}

// TestCompositeTypes tests that used composite types are returned.
func TestCompositeTypes(t *testing.T) {
	conn := setupTestDB(t)
	ctx := context.Background()

	typeName := "test_address_" + randomSuffix()
	_, err := conn.Exec(ctx, "CREATE TYPE "+typeName+" AS (street text, city text, zip varchar)")
	require.NoError(t, err)
	t.Cleanup(func() {
		conn.Exec(ctx, "DROP TYPE IF EXISTS "+typeName)
	})

	r := setupResolver(t)

	// First resolve a query that uses this composite type
	_, _, err = r.ResolveTypes(ctx, "SELECT ROW('Main St', 'NYC', '10001')::"+typeName, nil)
	require.NoError(t, err)

	// Now get composite types
	composites, err := r.CompositeTypes(ctx)
	require.NoError(t, err)

	goName := snakeToCamel(typeName)
	var found bool
	for _, c := range composites {
		if c.Name == goName {
			found = true
			assert.True(t, c.IsCompositeType)
			require.Len(t, c.Params, 3)
			break
		}
	}
	assert.True(t, found, "Used composite type should be returned")
}

// TestNullabilityFromTableColumn tests nullability detection from table columns.
func TestNullabilityFromTableColumn(t *testing.T) {
	conn := setupTestDB(t)
	ctx := context.Background()

	tableName := "test_nullability_" + randomSuffix()
	_, err := conn.Exec(ctx, "CREATE TABLE "+tableName+" (id int4 NOT NULL, name text, age int4 NOT NULL)")
	require.NoError(t, err)
	t.Cleanup(func() {
		conn.Exec(ctx, "DROP TABLE IF EXISTS "+tableName)
	})

	r := setupResolver(t)

	_, returns, err := r.ResolveTypes(ctx, "SELECT id, name, age FROM "+tableName, nil)
	require.NoError(t, err)
	require.NotNil(t, returns)
	require.Len(t, returns.Params, 3)

	// id is NOT NULL -> int
	assert.Equal(t, "int", returns.Params[0].Type.Name)
	// name is nullable -> *string
	assert.Equal(t, "*string", returns.Params[1].Type.Name)
	// age is NOT NULL -> int
	assert.Equal(t, "int", returns.Params[2].Type.Name)
}

// TestTimeTypes verifies time-related type mappings.
func TestTimeTypes(t *testing.T) {
	r := setupResolver(t)
	ctx := context.Background()

	tests := []struct {
		name     string
		pgType   string
		notNull  bool
		wantType string
	}{
		{"timestamp not null", "timestamp", true, "time.Time"},
		{"timestamp nullable", "timestamp", false, "pgtype.Timestamp"},
		{"timestamptz not null", "timestamptz", true, "time.Time"},
		{"timestamptz nullable", "timestamptz", false, "pgtype.Timestamptz"},
		{"date not null", "date", true, "time.Time"},
		{"date nullable", "date", false, "pgtype.Date"},
		{"interval not null", "interval", true, "pgtype.Interval"},
		{"interval nullable", "interval", false, "*pgtype.Interval"},
		{"time not null", "time", true, "pgtype.Time"},
		{"time nullable", "time", false, "pgtype.Time"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params, _, err := r.ResolveTypes(ctx, "SELECT $1::"+tt.pgType, []bool{tt.notNull})
			require.NoError(t, err)
			require.Len(t, params, 1)
			assert.Equal(t, tt.wantType, params[0].Name)
		})
	}
}

// TestNetworkTypes verifies network-related type mappings.
func TestNetworkTypes(t *testing.T) {
	r := setupResolver(t)
	ctx := context.Background()

	tests := []struct {
		name     string
		pgType   string
		notNull  bool
		wantType string
	}{
		{"inet not null", "inet", true, "netip.Addr"},
		{"inet nullable", "inet", false, "netip.Addr"},
		{"cidr not null", "cidr", true, "netip.Prefix"},
		{"cidr nullable", "cidr", false, "netip.Prefix"},
		{"macaddr not null", "macaddr", true, "net.HardwareAddr"},
		{"macaddr nullable", "macaddr", false, "net.HardwareAddr"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params, _, err := r.ResolveTypes(ctx, "SELECT $1::"+tt.pgType, []bool{tt.notNull})
			require.NoError(t, err)
			require.Len(t, params, 1)
			assert.Equal(t, tt.wantType, params[0].Name)
		})
	}
}

// TestDomainType tests domain type resolution.
// PostgreSQL returns base type OID for domain columns, but the domain's NOT NULL
// constraint should still be detected via pg_attribute and pg_type.
func TestDomainType(t *testing.T) {
	conn := setupTestDB(t)
	ctx := context.Background()

	domainName := "test_positive_int_" + randomSuffix()
	tableName := "test_domain_table_" + randomSuffix()
	_, err := conn.Exec(ctx, "CREATE DOMAIN "+domainName+" AS int4 NOT NULL CHECK (VALUE > 0)")
	require.NoError(t, err)
	_, err = conn.Exec(ctx, "CREATE TABLE "+tableName+" (val "+domainName+")")
	require.NoError(t, err)
	t.Cleanup(func() {
		conn.Exec(ctx, "DROP TABLE IF EXISTS "+tableName)
		conn.Exec(ctx, "DROP DOMAIN IF EXISTS "+domainName)
	})

	r := setupResolver(t)

	_, returns, err := r.ResolveTypes(ctx, "SELECT val FROM "+tableName, nil)
	require.NoError(t, err)
	require.NotNil(t, returns)
	require.Len(t, returns.Params, 1)
	// Domain's NOT NULL makes the column non-nullable
	assert.Equal(t, "int", returns.Params[0].Type.Name)
}

// TestNewlyAddedTypes tests the newly added type mappings.
func TestNewlyAddedTypes(t *testing.T) {
	r := setupResolver(t)
	ctx := context.Background()

	tests := []struct {
		name     string
		pgType   string
		notNull  bool
		wantType string
	}{
		{"bpchar not null", "bpchar", true, "string"},
		{"bpchar nullable", "bpchar", false, "*string"},
		// Note: ::char in PostgreSQL actually casts to bpchar, not internal "char" type
		{"char not null", "char", true, "string"},   // becomes bpchar
		{"char nullable", "char", false, "*string"}, // becomes bpchar
		{"name not null", "name", true, "string"},
		{"name nullable", "name", false, "*string"},
		{"varbit not null", "varbit", true, "pgtype.Bits"},
		{"varbit nullable", "varbit", false, "pgtype.Bits"},
		{"money not null", "money", true, "string"},
		{"money nullable", "money", false, "*string"},
		{"timetz not null", "timetz", true, "pgtype.Time"},
		{"timetz nullable", "timetz", false, "pgtype.Time"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params, _, err := r.ResolveTypes(ctx, "SELECT $1::"+tt.pgType, []bool{tt.notNull})
			require.NoError(t, err)
			require.Len(t, params, 1)
			assert.Equal(t, tt.wantType, params[0].Name)
		})
	}
}

// TestArrayTypes tests that PostgreSQL array types resolve to Go slices.
func TestArrayTypes(t *testing.T) {
	r := setupResolver(t)
	ctx := context.Background()

	tests := []struct {
		name     string
		query    string
		notnulls []bool
		wantType string
	}{
		{"uuid array", "SELECT $1::uuid[]", []bool{true}, "[]pgtype.UUID"},
		{"int4 array", "SELECT $1::int4[]", []bool{true}, "[]int"},
		{"text array", "SELECT $1::text[]", []bool{true}, "[]string"},
		{"bool array", "SELECT $1::bool[]", []bool{true}, "[]bool"},
		{"float8 array", "SELECT $1::float8[]", []bool{true}, "[]float64"},
		{"int8 array", "SELECT $1::int8[]", []bool{true}, "[]int"},
		{"timestamptz array", "SELECT $1::timestamptz[]", []bool{true}, "[]time.Time"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params, _, err := r.ResolveTypes(ctx, tt.query, tt.notnulls)
			require.NoError(t, err)
			require.Len(t, params, 1)
			assert.Equal(t, tt.wantType, params[0].Name)
		})
	}
}

// TestArrayTypeAsReturnColumn tests array types in SELECT return columns.
func TestArrayTypeAsReturnColumn(t *testing.T) {
	conn := setupTestDB(t)
	ctx := context.Background()

	tableName := "test_arrays_" + randomSuffix()
	_, err := conn.Exec(ctx, "CREATE TABLE "+tableName+" (id int4 NOT NULL, tags text[] NOT NULL, scores int4[])")
	require.NoError(t, err)
	t.Cleanup(func() {
		conn.Exec(ctx, "DROP TABLE IF EXISTS "+tableName)
	})

	r := setupResolver(t)

	_, returns, err := r.ResolveTypes(ctx, "SELECT id, tags, scores FROM "+tableName, nil)
	require.NoError(t, err)
	require.NotNil(t, returns)
	require.Len(t, returns.Params, 3)

	assert.Equal(t, "int", returns.Params[0].Type.Name)
	assert.Equal(t, "[]string", returns.Params[1].Type.Name)
	assert.Equal(t, "[]int", returns.Params[2].Type.Name)
}

// TestArrayTypeWithAny tests the common pattern of using array param with ANY().
func TestArrayTypeWithAny(t *testing.T) {
	conn := setupTestDB(t)
	ctx := context.Background()

	tableName := "test_any_arr_" + randomSuffix()
	_, err := conn.Exec(ctx, "CREATE TABLE "+tableName+" (id uuid NOT NULL DEFAULT gen_random_uuid(), name text NOT NULL)")
	require.NoError(t, err)
	t.Cleanup(func() {
		conn.Exec(ctx, "DROP TABLE IF EXISTS "+tableName)
	})

	r := setupResolver(t)

	// This is the pattern from the original bug report: id = any($1::uuid[])
	params, returns, err := r.ResolveTypes(ctx, "SELECT id, name FROM "+tableName+" WHERE id = any($1::uuid[])", []bool{true})
	require.NoError(t, err)
	require.Len(t, params, 1)
	assert.Equal(t, "[]pgtype.UUID", params[0].Name)
	require.NotNil(t, returns)
	require.Len(t, returns.Params, 2)
}

// TestUnknownType tests error handling for types without mappings.
func TestUnknownType(t *testing.T) {
	ctx := context.Background()
	// Create resolver with no type mappings
	r, err := NewResolver(ctx, getTestConnString(t), nil)
	require.NoError(t, err)
	defer r.Close()

	_, _, err = r.ResolveTypes(ctx, "SELECT $1::int4", []bool{true})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unknown type")
}

// TestMultipleParams tests resolution of multiple parameters.
func TestMultipleParams(t *testing.T) {
	r := setupResolver(t)
	ctx := context.Background()

	params, _, err := r.ResolveTypes(ctx, "SELECT $1::int4 + $2::int4 + $3::int8", []bool{true, false, true})
	require.NoError(t, err)
	require.Len(t, params, 3)
	assert.Equal(t, "int", params[0].Name)
	assert.Equal(t, "*int", params[1].Name)
	assert.Equal(t, "int", params[2].Name)
}

// TestEnumValuesOrder tests that enum values maintain their sort order.
func TestEnumValuesOrder(t *testing.T) {
	conn := setupTestDB(t)
	ctx := context.Background()

	enumName := "test_ordered_" + randomSuffix()
	_, err := conn.Exec(ctx, "CREATE TYPE "+enumName+" AS ENUM ('zebra', 'apple', 'mango')")
	require.NoError(t, err)
	t.Cleanup(func() {
		conn.Exec(ctx, "DROP TYPE IF EXISTS "+enumName)
	})

	r := setupResolver(t)
	enums := r.Enums()

	for _, e := range enums {
		if e.Name == snakeToCamel(enumName) {
			// Values should be in definition order, not alphabetical
			assert.Equal(t, []string{"zebra", "apple", "mango"}, e.Values)
			return
		}
	}
	t.Error("Created enum not found")
}

// TestCompositeTypeFieldNullability tests field nullability in composite types.
func TestCompositeTypeFieldNullability(t *testing.T) {
	conn := setupTestDB(t)
	ctx := context.Background()

	typeName := "test_mixed_nulls_" + randomSuffix()
	_, err := conn.Exec(ctx, "CREATE TYPE "+typeName+" AS (required int4, optional text)")
	require.NoError(t, err)
	t.Cleanup(func() {
		conn.Exec(ctx, "DROP TYPE IF EXISTS "+typeName)
	})

	r := setupResolver(t)

	_, returns, err := r.ResolveTypes(ctx, "SELECT ROW(1, 'x')::"+typeName, nil)
	require.NoError(t, err)
	require.NotNil(t, returns)
	require.Len(t, returns.Params, 2)

	// Composite type fields are nullable by default (no column constraint)
	assert.Equal(t, "*int", returns.Params[0].Type.Name)
	assert.Equal(t, "*string", returns.Params[1].Type.Name)
}

// TestEnumsDoesNotPanic tests that Enums() doesn't panic.
func TestEnumsDoesNotPanic(t *testing.T) {
	r := setupResolver(t)
	// Should not panic; may return nil or empty slice
	_ = r.Enums()
}

// TestCompositeTypesNotUsed tests that unused composite types are not returned.
func TestCompositeTypesNotUsed(t *testing.T) {
	conn := setupTestDB(t)
	ctx := context.Background()

	typeName := "test_unused_" + randomSuffix()
	_, err := conn.Exec(ctx, "CREATE TYPE "+typeName+" AS (field1 int4)")
	require.NoError(t, err)
	t.Cleanup(func() {
		conn.Exec(ctx, "DROP TYPE IF EXISTS "+typeName)
	})

	r := setupResolver(t)

	// Don't use the type in any query
	composites, err := r.CompositeTypes(ctx)
	require.NoError(t, err)

	goName := snakeToCamel(typeName)
	for _, c := range composites {
		assert.NotEqual(t, goName, c.Name, "Unused composite type should not be returned")
	}
}

// randomSuffix generates a unique suffix for test object names using timestamp.
func randomSuffix() string {
	return fmt.Sprintf("%d", time.Now().UnixNano()%1000000)
}

// snakeToCamel is a simple helper that mimics the snaker behavior for test assertions.
func snakeToCamel(s string) string {
	var result strings.Builder
	capitalizeNext := true
	for _, c := range s {
		if c == '_' {
			capitalizeNext = true
			continue
		}
		if capitalizeNext && c >= 'a' && c <= 'z' {
			result.WriteByte(byte(c - 32)) // ASCII uppercase
			capitalizeNext = false
		} else {
			result.WriteRune(c)
			capitalizeNext = false
		}
	}
	return result.String()
}
