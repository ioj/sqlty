package generator

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ioj/sqlty/stmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupGenerator creates a generator with temp directories for testing.
func setupGenerator(t *testing.T) (*Generator, string) {
	t.Helper()
	outDir := t.TempDir()
	cacheDir := t.TempDir()

	g, err := New("", cacheDir)
	require.NoError(t, err)
	t.Cleanup(func() { g.Close() })

	return g, outDir
}

// readGeneratedFile reads and returns the content of a generated file.
func readGeneratedFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	require.NoError(t, err)
	return string(data)
}

// intParam creates a simple int parameter for tests.
func intParam(name string) stmt.Param {
	return stmt.Param{Name: name, Type: stmt.Type{Name: "int", ZeroValue: "0"}}
}

// stringParam creates a simple string parameter for tests.
func stringParam(name string) stmt.Param {
	return stmt.Param{Name: name, Type: stmt.Type{Name: "string", ZeroValue: `""`}}
}

func TestNew(t *testing.T) {
	cacheDir := t.TempDir()
	g, err := New("", cacheDir)
	require.NoError(t, err)
	require.NotNil(t, g)
	require.NoError(t, g.Close())
}

func TestNew_InvalidTemplateDir(t *testing.T) {
	_, err := New("/nonexistent/path/templates", t.TempDir())
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse custom templates")
}

func TestDB(t *testing.T) {
	g, outDir := setupGenerator(t)

	err := g.DB(outDir, &stmt.DB{PackageName: "testpkg"})
	require.NoError(t, err)

	content := readGeneratedFile(t, filepath.Join(outDir, "db.sqlty.gen.go"))
	assert.Contains(t, content, "package testpkg")
	assert.Contains(t, content, "type DB struct")
	assert.Contains(t, content, "func New(db Txer)")
	assert.Contains(t, content, "Middleware")
	assert.Contains(t, content, "CtxTx")
}

func TestDB_NilInput(t *testing.T) {
	g, outDir := setupGenerator(t)
	err := g.DB(outDir, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "db is required")
}

func TestEnums(t *testing.T) {
	g, outDir := setupGenerator(t)

	enums := &stmt.Enums{
		PackageName: "testpkg",
		Enums:       []*stmt.Enum{{Name: "Status", Values: []string{"active", "pending", "done"}}},
	}

	err := g.Enums(outDir, enums)
	require.NoError(t, err)

	content := readGeneratedFile(t, filepath.Join(outDir, "enums.sqlty.gen.go"))
	assert.Contains(t, content, "package testpkg")
	assert.Contains(t, content, "type Status string")
	assert.Contains(t, content, `Status_Active Status = "active"`)
	assert.Contains(t, content, "StatusIdxMap")
}

func TestEnums_Empty(t *testing.T) {
	g, outDir := setupGenerator(t)

	fname := filepath.Join(outDir, "enums.sqlty.gen.go")
	err := os.WriteFile(fname, []byte("dummy"), 0644)
	require.NoError(t, err)

	err = g.Enums(outDir, &stmt.Enums{PackageName: "testpkg", Enums: nil})
	require.NoError(t, err)

	_, err = os.Stat(fname)
	assert.True(t, os.IsNotExist(err))
}

func TestEnums_Nil(t *testing.T) {
	g, outDir := setupGenerator(t)
	err := g.Enums(outDir, nil)
	require.NoError(t, err)
}

func TestCompositeTypes(t *testing.T) {
	g, outDir := setupGenerator(t)

	types := &stmt.CompositeTypes{
		PackageName: "testpkg",
		Types: []*stmt.Struct{{
			Name:            "Address",
			IsCompositeType: true,
			Params:          []stmt.Param{stringParam("Street"), stringParam("City")},
		}},
	}

	err := g.CompositeTypes(outDir, types)
	require.NoError(t, err)

	content := readGeneratedFile(t, filepath.Join(outDir, "composite_types.sqlty.gen.go"))
	assert.Contains(t, content, "package testpkg")
	assert.Contains(t, content, "type Address struct")
	assert.Contains(t, content, `json:"street"`)
}

func TestCompositeTypes_Empty(t *testing.T) {
	g, outDir := setupGenerator(t)

	fname := filepath.Join(outDir, "composite_types.sqlty.gen.go")
	err := os.WriteFile(fname, []byte("dummy"), 0644)
	require.NoError(t, err)

	err = g.CompositeTypes(outDir, &stmt.CompositeTypes{PackageName: "testpkg", Types: nil})
	require.NoError(t, err)

	_, err = os.Stat(fname)
	assert.True(t, os.IsNotExist(err))
}

func TestQuery_NilInput(t *testing.T) {
	g, outDir := setupGenerator(t)
	err := g.Query("array", filepath.Join(outDir, "test.gen.go"), nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "query is required")
}

func TestQuery_InvalidTemplate(t *testing.T) {
	g, outDir := setupGenerator(t)
	q := &stmt.Query{
		PackageName: "testpkg",
		Name:        "GetUser",
		Statement:   "SELECT id FROM users",
		ExecMode:    stmt.ExecModeOne,
		Returns:     stmt.Struct{Params: []stmt.Param{intParam("Id")}},
	}

	err := g.Query("nonexistent", filepath.Join(outDir, "test.gen.go"), q)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "template not found")
}

// TestQuery_ExecutionModes tests all three execution modes in a table-driven test.
func TestQuery_ExecutionModes(t *testing.T) {
	tests := []struct {
		name     string
		execMode stmt.ExecMode
		returns  stmt.Struct
		contains []string
	}{
		{
			name:     "exec mode",
			execMode: stmt.ExecModeExec,
			contains: []string{"pgconn.CommandTag", "db.tx.Exec"},
		},
		{
			name:     "one mode",
			execMode: stmt.ExecModeOne,
			returns:  stmt.Struct{Params: []stmt.Param{intParam("Id")}},
			contains: []string{"CollectOneRow", "RowTo[int]"},
		},
		{
			name:     "many mode",
			execMode: stmt.ExecModeMany,
			returns:  stmt.Struct{Params: []stmt.Param{intParam("Id")}},
			contains: []string{"CollectRows", "[]int"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, outDir := setupGenerator(t)
			q := &stmt.Query{
				PackageName: "testpkg",
				Name:        "TestQuery",
				Statement:   "SELECT id FROM users WHERE id = $1",
				ExecMode:    tt.execMode,
				Params:      stmt.Params{Scalar: []stmt.Param{intParam("id")}},
				Returns:     tt.returns,
			}

			fname := filepath.Join(outDir, "test.gen.go")
			err := g.Query("array", fname, q)
			require.NoError(t, err)

			content := readGeneratedFile(t, fname)
			for _, s := range tt.contains {
				assert.Contains(t, content, s)
			}
		})
	}
}

func TestQuery_WithStructReturn(t *testing.T) {
	g, outDir := setupGenerator(t)

	q := &stmt.Query{
		PackageName: "testpkg",
		Name:        "GetUserDetails",
		Statement:   "SELECT id, name FROM users WHERE id = $1",
		ExecMode:    stmt.ExecModeOne,
		Params:      stmt.Params{Scalar: []stmt.Param{intParam("id")}},
		Returns: stmt.Struct{
			Name:   "GetUserDetailsRow",
			Params: []stmt.Param{intParam("Id"), stringParam("Name")},
		},
	}

	fname := filepath.Join(outDir, "get_user_details.gen.go")
	err := g.Query("array", fname, q)
	require.NoError(t, err)

	content := readGeneratedFile(t, fname)
	assert.Contains(t, content, "type GetUserDetailsRow struct")
	assert.Contains(t, content, "*GetUserDetailsRow")
	assert.Contains(t, content, "RowToAddrOfStructByPos")
}

func TestQuery_WithCompositeTypeReturn(t *testing.T) {
	g, outDir := setupGenerator(t)

	q := &stmt.Query{
		PackageName: "testpkg",
		Name:        "GetAddress",
		Statement:   "SELECT address FROM users WHERE id = $1",
		ExecMode:    stmt.ExecModeOne,
		Params:      stmt.Params{Scalar: []stmt.Param{intParam("id")}},
		Returns: stmt.Struct{
			Name:            "Address",
			IsCompositeType: true,
			Params:          []stmt.Param{stringParam("Street"), stringParam("City")},
		},
	}

	fname := filepath.Join(outDir, "get_address.gen.go")
	err := g.Query("array", fname, q)
	require.NoError(t, err)

	content := readGeneratedFile(t, fname)
	assert.Contains(t, content, "*Address")
	assert.Contains(t, content, "RowToAddrOfStructByPos[Address]")
	// Composite types should NOT generate inline struct declaration
	assert.NotContains(t, content, "type Address struct")
}

func TestQuery_WithSpreadParams(t *testing.T) {
	g, outDir := setupGenerator(t)

	q := &stmt.Query{
		PackageName: "testpkg",
		Name:        "GetUsersByIDs",
		Statement:   "SELECT id FROM users WHERE id IN (%s)",
		ExecMode:    stmt.ExecModeMany,
		Params:      stmt.Params{Spread: []stmt.Param{intParam("ids")}},
		Returns:     stmt.Struct{Params: []stmt.Param{intParam("Id")}},
	}

	fname := filepath.Join(outDir, "get_users_by_ids.gen.go")
	err := g.Query("array", fname, q)
	require.NoError(t, err)

	content := readGeneratedFile(t, fname)
	assert.Contains(t, content, "ids []int")
	assert.Contains(t, content, "fmt.Sprintf")
	assert.Contains(t, content, "ErrEmptySpread")
}

func TestQuery_WithStructSpread(t *testing.T) {
	g, outDir := setupGenerator(t)

	q := &stmt.Query{
		PackageName: "testpkg",
		Name:        "InsertUsers",
		Statement:   "INSERT INTO users (name, age) VALUES %s",
		ExecMode:    stmt.ExecModeExec,
		Params: stmt.Params{
			StructSpread: []stmt.Struct{{
				Name:   "User",
				Params: []stmt.Param{stringParam("Name"), intParam("Age")},
			}},
		},
	}

	fname := filepath.Join(outDir, "insert_users.gen.go")
	err := g.Query("array", fname, q)
	require.NoError(t, err)

	content := readGeneratedFile(t, fname)
	assert.Contains(t, content, "type User struct")
	assert.Contains(t, content, "User []*User")
	assert.Contains(t, content, "ErrEmptyStructSpreadList")
}

func TestQuery_ExecModeWithSpreadParams(t *testing.T) {
	g, outDir := setupGenerator(t)

	q := &stmt.Query{
		PackageName: "testpkg",
		Name:        "DeleteUsers",
		Statement:   "DELETE FROM users WHERE id IN (%s)",
		ExecMode:    stmt.ExecModeExec,
		Params:      stmt.Params{Spread: []stmt.Param{intParam("ids")}},
	}

	fname := filepath.Join(outDir, "delete_users.gen.go")
	err := g.Query("array", fname, q)
	require.NoError(t, err)

	content := readGeneratedFile(t, fname)
	assert.Contains(t, content, "pgconn.CommandTag")
	assert.Contains(t, content, "ErrEmptySpread")
	assert.Contains(t, content, "return pgconn.CommandTag{}, ErrEmptySpread")
}

func TestQuery_ExecModeWithParamsStruct(t *testing.T) {
	g, outDir := setupGenerator(t)

	q := &stmt.Query{
		PackageName: "testpkg",
		Name:        "UpdateUser",
		Statement:   "UPDATE users SET name = $1 WHERE id = $2",
		ExecMode:    stmt.ExecModeExec,
		Params: stmt.Params{
			Name:   "UpdateUserParams",
			Scalar: []stmt.Param{stringParam("Name"), intParam("Id")},
		},
	}

	fname := filepath.Join(outDir, "update_user.gen.go")
	err := g.Query("array", fname, q)
	require.NoError(t, err)

	content := readGeneratedFile(t, fname)
	assert.Contains(t, content, "type UpdateUserParams struct")
	assert.Contains(t, content, "args *UpdateUserParams")
	assert.Contains(t, content, "ErrNilArgs")
	assert.Contains(t, content, "return pgconn.CommandTag{}, ErrNilArgs")
}

func TestQuery_WithComments(t *testing.T) {
	g, outDir := setupGenerator(t)

	q := &stmt.Query{
		PackageName: "testpkg",
		Name:        "GetUser",
		Statement:   "SELECT id FROM users",
		ExecMode:    stmt.ExecModeOne,
		Comments:    []string{"GetUser retrieves a user by ID.", "Returns ErrNoRows if not found."},
		Returns:     stmt.Struct{Params: []stmt.Param{intParam("Id")}},
	}

	fname := filepath.Join(outDir, "get_user.gen.go")
	err := g.Query("array", fname, q)
	require.NoError(t, err)

	content := readGeneratedFile(t, fname)
	assert.Contains(t, content, "// GetUser retrieves a user by ID.")
	assert.Contains(t, content, "// Returns ErrNoRows if not found.")
}

func TestQuery_NoParams(t *testing.T) {
	g, outDir := setupGenerator(t)

	q := &stmt.Query{
		PackageName: "testpkg",
		Name:        "ListAllUsers",
		Statement:   "SELECT id FROM users",
		ExecMode:    stmt.ExecModeMany,
		Returns:     stmt.Struct{Params: []stmt.Param{intParam("Id")}},
	}

	fname := filepath.Join(outDir, "list_all_users.gen.go")
	err := g.Query("array", fname, q)
	require.NoError(t, err)

	content := readGeneratedFile(t, fname)
	assert.Contains(t, content, "func (db *DB) ListAllUsers(ctx context.Context)")
	assert.NotContains(t, content, "sqltyStmtargs")
}

func TestGenerator_CacheSkipsUnchanged(t *testing.T) {
	cacheDir := t.TempDir()
	outDir := t.TempDir()

	q := &stmt.Query{
		PackageName: "testpkg",
		Name:        "GetUser",
		Statement:   "SELECT id FROM users",
		ExecMode:    stmt.ExecModeOne,
		Returns:     stmt.Struct{Params: []stmt.Param{intParam("Id")}},
	}

	fname := filepath.Join(outDir, "test.gen.go")

	// First generation
	g1, err := New("", cacheDir)
	require.NoError(t, err)
	err = g1.Query("array", fname, q)
	require.NoError(t, err)
	require.NoError(t, g1.Close())

	info1, err := os.Stat(fname)
	require.NoError(t, err)

	// Second generation with same params
	g2, err := New("", cacheDir)
	require.NoError(t, err)
	err = g2.Query("array", fname, q)
	require.NoError(t, err)
	require.NoError(t, g2.Close())

	info2, err := os.Stat(fname)
	require.NoError(t, err)
	assert.Equal(t, info1.ModTime(), info2.ModTime())
}

func TestGenerator_CacheRegeneratesOnChange(t *testing.T) {
	cacheDir := t.TempDir()
	outDir := t.TempDir()

	q := &stmt.Query{
		PackageName: "testpkg",
		Name:        "GetUser",
		Statement:   "SELECT id FROM users",
		ExecMode:    stmt.ExecModeOne,
		Returns:     stmt.Struct{Params: []stmt.Param{intParam("Id")}},
	}

	fname := filepath.Join(outDir, "test.gen.go")

	// First generation
	g1, err := New("", cacheDir)
	require.NoError(t, err)
	err = g1.Query("array", fname, q)
	require.NoError(t, err)
	require.NoError(t, g1.Close())

	content1 := readGeneratedFile(t, fname)

	// Second generation with different statement
	q.Statement = "SELECT id FROM users WHERE active = true"
	g2, err := New("", cacheDir)
	require.NoError(t, err)
	err = g2.Query("array", fname, q)
	require.NoError(t, err)
	require.NoError(t, g2.Close())

	content2 := readGeneratedFile(t, fname)
	assert.NotEqual(t, content1, content2)
	assert.Contains(t, content2, "WHERE active = true")
}

// TestGeneratedCodeCompiles verifies that generated code actually compiles.
func TestGeneratedCodeCompiles(t *testing.T) {
	g, outDir := setupGenerator(t)

	// Generate all types of files
	require.NoError(t, g.DB(outDir, &stmt.DB{PackageName: "testpkg"}))

	require.NoError(t, g.Enums(outDir, &stmt.Enums{
		PackageName: "testpkg",
		Enums:       []*stmt.Enum{{Name: "Status", Values: []string{"active", "pending"}}},
	}))

	require.NoError(t, g.CompositeTypes(outDir, &stmt.CompositeTypes{
		PackageName: "testpkg",
		Types: []*stmt.Struct{{
			Name:            "Address",
			IsCompositeType: true,
			Params:          []stmt.Param{stringParam("City")},
		}},
	}))

	// Generate queries covering different patterns
	queries := []*stmt.Query{
		{
			PackageName: "testpkg",
			Name:        "GetUser",
			Statement:   "SELECT id FROM users WHERE id = $1",
			ExecMode:    stmt.ExecModeOne,
			Params:      stmt.Params{Scalar: []stmt.Param{intParam("id")}},
			Returns:     stmt.Struct{Params: []stmt.Param{intParam("Id")}},
		},
		{
			PackageName: "testpkg",
			Name:        "ListUsers",
			Statement:   "SELECT id, name FROM users",
			ExecMode:    stmt.ExecModeMany,
			Returns: stmt.Struct{
				Name:   "ListUsersRow",
				Params: []stmt.Param{intParam("Id"), stringParam("Name")},
			},
		},
		{
			PackageName: "testpkg",
			Name:        "DeleteUser",
			Statement:   "DELETE FROM users WHERE id = $1",
			ExecMode:    stmt.ExecModeExec,
			Params:      stmt.Params{Scalar: []stmt.Param{intParam("id")}},
		},
		{
			PackageName: "testpkg",
			Name:        "DeleteUsers",
			Statement:   "DELETE FROM users WHERE id IN (%s)",
			ExecMode:    stmt.ExecModeExec,
			Params:      stmt.Params{Spread: []stmt.Param{intParam("ids")}},
		},
	}

	for _, q := range queries {
		fname := filepath.Join(outDir, strings.ToLower(q.Name)+".gen.go")
		require.NoError(t, g.Query("array", fname, q))
	}

	// Create go.mod
	goMod := `module testpkg

go 1.21

require github.com/jackc/pgx/v5 v5.5.0
`
	require.NoError(t, os.WriteFile(filepath.Join(outDir, "go.mod"), []byte(goMod), 0644))

	// Run goimports
	cmd := exec.Command("goimports", "-w", outDir)
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Logf("goimports: %s", output)
		t.Skip("goimports not available")
	}

	// Run go mod tidy
	cmd = exec.Command("go", "mod", "tidy")
	cmd.Dir = outDir
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("go mod tidy failed: %s\n%s", err, output)
	}

	// Run go build
	cmd = exec.Command("go", "build", "./...")
	cmd.Dir = outDir
	if output, err := cmd.CombinedOutput(); err != nil {
		files, _ := filepath.Glob(filepath.Join(outDir, "*.go"))
		for _, f := range files {
			content, _ := os.ReadFile(f)
			t.Logf("=== %s ===\n%s\n", f, content)
		}
		t.Fatalf("go build failed: %s\n%s", err, output)
	}
}
