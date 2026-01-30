package compiler

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCompileString_SimpleSelectAll(t *testing.T) {
	input := `
/*
	@name GetAllUsers
	@many
*/
SELECT * FROM users;`

	query, err := CompileString("test.sql", input)

	require.NoError(t, err)
	require.NotNil(t, query)

	assert.Equal(t, "GetAllUsers", query.Name())
	assert.Equal(t, "SELECT * FROM users", query.Statement())
	assert.Equal(t, "array", query.Template())
	assert.Equal(t, "test.sql", query.Filename)

	assert.Empty(t, query.Params(Scalar))
	assert.Empty(t, query.Params(Spread))
	assert.Empty(t, query.Params(StructSpread))

	assert.False(t, query.NeedsSprintf())
	assert.Empty(t, query.NotNullArray())
	assert.Empty(t, query.Comments)
}

func TestCompileString_MultiLineAnnotation(t *testing.T) {
	input := `/*
  @name GetAllUsers
	@many
*/
SELECT id, name FROM users;`

	query, err := CompileString("test.sql", input)

	require.NoError(t, err)
	require.NotNil(t, query)

	assert.Equal(t, "GetAllUsers", query.Name())
	assert.Equal(t, "SELECT id, name FROM users", query.Statement())
	assert.Equal(t, "array", query.Template())
	assert.Equal(t, "test.sql", query.Filename)

	assert.Empty(t, query.Params(Scalar))
	assert.Empty(t, query.Params(Spread))
	assert.Empty(t, query.Params(StructSpread))

	assert.False(t, query.NeedsSprintf())
	assert.Empty(t, query.NotNullArray())
	assert.Empty(t, query.Comments)
}

func TestCompileString_WithScalarParam(t *testing.T) {
	input := ` /* @name GetUserById @many */
  SELECT * FROM users WHERE userId = :userId;`

	query, err := CompileString("test.sql", input)

	require.NoError(t, err)
	require.NotNil(t, query)

	assert.Equal(t, "GetUserById", query.Name())
	assert.Equal(t, "SELECT * FROM users WHERE userId = $1", query.Statement())
	assert.Equal(t, "array", query.Template())
	assert.Equal(t, "test.sql", query.Filename)

	scalars := query.Params(Scalar)
	require.Len(t, scalars, 1)
	assert.Equal(t, Scalar, scalars[0].Type)
	assert.Equal(t, 0, scalars[0].Idx)
	assert.False(t, scalars[0].NotNull)

	assert.Empty(t, query.Params(Spread))
	assert.Empty(t, query.Params(StructSpread))

	assert.False(t, query.NeedsSprintf())
	assert.Equal(t, []bool{false}, query.NotNullArray())
	assert.Empty(t, query.Comments)
}

func TestCompileString_WithReusedScalarParam(t *testing.T) {
	input := `/* @name GetUserById @many */
  SELECT * FROM users WHERE userId = :userId or parentId = :userId;`

	query, err := CompileString("test.sql", input)

	require.NoError(t, err)
	require.NotNil(t, query)

	assert.Equal(t, "GetUserById", query.Name())
	assert.Equal(t, "SELECT * FROM users WHERE userId = $1 or parentId = $1", query.Statement())
	assert.Equal(t, "array", query.Template())
	assert.Equal(t, "test.sql", query.Filename)

	scalars := query.Params(Scalar)
	require.Len(t, scalars, 1)
	assert.Equal(t, Scalar, scalars[0].Type)
	assert.Equal(t, 0, scalars[0].Idx)
	assert.False(t, scalars[0].NotNull)

	assert.Empty(t, query.Params(Spread))
	assert.Empty(t, query.Params(StructSpread))

	assert.False(t, query.NeedsSprintf())
	assert.Equal(t, []bool{false}, query.NotNullArray())
	assert.Empty(t, query.Comments)
}

func TestCompileString_WithStructSpreadParam(t *testing.T) {
	input := `/*
    @exec
    @name CreateCustomer
    @param customers -> ((customerName, contactName, address)...)
*/
  INSERT INTO customers (customer_name, contact_name, address)
  VALUES :customers;`

	query, err := CompileString("test.sql", input)

	require.NoError(t, err)
	require.NotNil(t, query)

	assert.Equal(t, "CreateCustomer", query.Name())
	assert.Equal(t, "array", query.Template())
	assert.Equal(t, "test.sql", query.Filename)

	assert.Empty(t, query.Params(Scalar))
	assert.Empty(t, query.Params(Spread))

	structSpreads := query.Params(StructSpread)
	require.Len(t, structSpreads, 1)
	assert.Equal(t, StructSpread, structSpreads[0].Type)
	assert.Equal(t, 0, structSpreads[0].Idx)
	assert.False(t, structSpreads[0].NotNull)

	keys := structSpreads[0].Keys()
	require.Len(t, keys, 3)
	assert.Equal(t, "customerName", keys[0].Name())
	assert.Equal(t, "contactName", keys[1].Name())
	assert.Equal(t, "address", keys[2].Name())

	assert.True(t, query.NeedsSprintf())
	assert.Equal(t, []bool{false, false, false}, query.NotNullArray())
	assert.Empty(t, query.Comments)
}

func TestCompileString_WithPostgresCastOperator(t *testing.T) {
	input := `/* @name GetAllUsers @many */
SELECT u."rank" FROM users u where name = :name::text;`

	query, err := CompileString("test.sql", input)

	require.NoError(t, err)
	require.NotNil(t, query)

	assert.Equal(t, "GetAllUsers", query.Name())
	assert.Equal(t, `SELECT u."rank" FROM users u where name = $1::text`, query.Statement())
	assert.Equal(t, "array", query.Template())
	assert.Equal(t, "test.sql", query.Filename)

	scalars := query.Params(Scalar)
	require.Len(t, scalars, 1)
	assert.Equal(t, Scalar, scalars[0].Type)
	assert.Equal(t, 0, scalars[0].Idx)
	assert.False(t, scalars[0].NotNull)

	assert.Empty(t, query.Params(Spread))
	assert.Empty(t, query.Params(StructSpread))

	assert.False(t, query.NeedsSprintf())
	assert.Equal(t, []bool{false}, query.NotNullArray())
	assert.Empty(t, query.Comments)
}
