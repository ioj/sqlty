package compiler

import (
	_ "embed"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

//go:embed tests/test_cases.yaml
var testCasesYAML []byte

// StructKeySnapshot represents expected values for a StructKey.
type StructKeySnapshot struct {
	Name    string `yaml:"name"`
	NotNull bool   `yaml:"notNull"`
}

// ParamSnapshot represents expected values for a Param.
type ParamSnapshot struct {
	Type    ParamType           `yaml:"type"`
	Idx     int                 `yaml:"idx"`
	NotNull bool                `yaml:"notNull"`
	Keys    []StructKeySnapshot `yaml:"keys"`
}

// QuerySnapshot represents expected values for a compiled Query.
type QuerySnapshot struct {
	Name          string          `yaml:"name"`
	Statement     string          `yaml:"statement"`
	Template      string          `yaml:"template"`
	Scalars       []ParamSnapshot `yaml:"scalars"`
	Spreads       []ParamSnapshot `yaml:"spreads"`
	StructSpreads []ParamSnapshot `yaml:"structSpreads"`
	NeedsSprintf  bool            `yaml:"needsSprintf"`
	NotNullArray  []bool          `yaml:"notNullArray"`
	Comments      []string        `yaml:"comments"`
}

// TestUseCase represents a single test case loaded from YAML.
type TestUseCase struct {
	Name          string        `yaml:"name"`
	Input         string        `yaml:"input"`
	Expected      QuerySnapshot `yaml:"expected"`
	ExpectedError string        `yaml:"expectedError"`
}

// assertParamEquals compares a single Param against an expected ParamSnapshot.
func assertParamEquals(t *testing.T, param *Param, expected ParamSnapshot, label string) {
	t.Helper()

	assert.Equal(t, expected.Type, param.Type, "%s.Type mismatch", label)
	assert.Equal(t, expected.Idx, param.Idx, "%s.Idx mismatch", label)
	assert.Equal(t, expected.NotNull, param.NotNull, "%s.NotNull mismatch", label)

	keys := param.Keys()
	require.Len(t, keys, len(expected.Keys), "%s.Keys count mismatch", label)
	for j, expKey := range expected.Keys {
		assert.Equal(t, expKey.Name, keys[j].Name(), "%s.Keys[%d].Name mismatch", label, j)
		assert.Equal(t, expKey.NotNull, keys[j].NotNull, "%s.Keys[%d].NotNull mismatch", label, j)
	}
}

// assertParamsEquals compares actual Params against expected ParamSnapshots.
func assertParamsEquals(t *testing.T, params []*Param, expected []ParamSnapshot, paramType string) {
	t.Helper()

	require.Len(t, params, len(expected), "%s params count mismatch", paramType)
	for i, exp := range expected {
		assertParamEquals(t, params[i], exp, fmt.Sprintf("%s[%d]", paramType, i))
	}
}

// assertQueryEquals compares a compiled Query against an expected QuerySnapshot.
func assertQueryEquals(t *testing.T, query *Query, expected QuerySnapshot) {
	t.Helper()

	assert.Equal(t, expected.Name, query.Name(), "Name mismatch")
	assert.Equal(t, strings.TrimSuffix(expected.Statement, "\n"), query.Statement(), "Statement mismatch")
	assert.Equal(t, expected.Template, query.Template(), "Template mismatch")

	assertParamsEquals(t, query.Params(Scalar), expected.Scalars, "Scalar")
	assertParamsEquals(t, query.Params(Spread), expected.Spreads, "Spread")
	assertParamsEquals(t, query.Params(StructSpread), expected.StructSpreads, "StructSpread")

	assert.Equal(t, expected.NeedsSprintf, query.NeedsSprintf(), "NeedsSprintf mismatch")
	assert.Equal(t, expected.NotNullArray, query.NotNullArray(), "NotNullArray mismatch")
	assert.Equal(t, expected.Comments, query.Comments, "Comments mismatch")
}

func TestCompileString(t *testing.T) {
	var testCases []TestUseCase
	err := yaml.Unmarshal(testCasesYAML, &testCases)
	require.NoError(t, err, "Failed to parse test cases YAML")

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			query, err := CompileString("test.sql", tc.Input)

			if tc.ExpectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.ExpectedError)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, query)

			// Set default template if not specified
			expected := tc.Expected
			if expected.Template == "" {
				expected.Template = "array"
			}

			assertQueryEquals(t, query, expected)
		})
	}
}
