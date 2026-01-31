package generator

import (
	"testing"

	"github.com/ioj/sqlty/stmt"
	"github.com/stretchr/testify/assert"
)

func TestFirstParamTypeName(t *testing.T) {
	fn := tmplfn["firstParamTypeName"].(func([]stmt.Param) string)

	tests := []struct {
		name   string
		params []stmt.Param
		want   string
	}{
		{"empty", nil, ""},
		{"single", []stmt.Param{{Name: "id", Type: stmt.Type{Name: "int"}}}, "int"},
		{"pointer", []stmt.Param{{Name: "name", Type: stmt.Type{Name: "*string"}}}, "*string"},
		{"multiple", []stmt.Param{
			{Name: "id", Type: stmt.Type{Name: "int"}},
			{Name: "name", Type: stmt.Type{Name: "string"}},
		}, "int"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, fn(tt.params))
		})
	}
}

func TestFirstParamNilReturnValue(t *testing.T) {
	fn := tmplfn["firstParamNilReturnValue"].(func([]stmt.Param) string)

	tests := []struct {
		name   string
		params []stmt.Param
		want   string
	}{
		{"empty", nil, ""},
		{"nullable", []stmt.Param{{Type: stmt.Type{Nullable: true, ZeroValue: `""`}}}, "nil"},
		{"not nullable", []stmt.Param{{Type: stmt.Type{Nullable: false, ZeroValue: `""`}}}, `""`},
		{"int zero", []stmt.Param{{Type: stmt.Type{Nullable: false, ZeroValue: "0"}}}, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, fn(tt.params))
		})
	}
}

func TestFirstParamZeroReturnValue(t *testing.T) {
	fn := tmplfn["firstParamZeroReturnValue"].(func([]stmt.Param) string)

	tests := []struct {
		name   string
		params []stmt.Param
		want   string
	}{
		{"empty", nil, ""},
		{"string", []stmt.Param{{Type: stmt.Type{ZeroValue: `""`}}}, `""`},
		{"int", []stmt.Param{{Type: stmt.Type{ZeroValue: "0"}}}, "0"},
		{"struct", []stmt.Param{{Type: stmt.Type{ZeroValue: "&User{}"}}}, "&User{}"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, fn(tt.params))
		})
	}
}

func TestNeedsPrintf(t *testing.T) {
	fn := tmplfn["needsPrintf"].(func(*stmt.Query) bool)

	tests := []struct {
		name  string
		query *stmt.Query
		want  bool
	}{
		{"no params", &stmt.Query{}, false},
		{"scalar only", &stmt.Query{Params: stmt.Params{Scalar: []stmt.Param{{Name: "id"}}}}, false},
		{"with spread", &stmt.Query{Params: stmt.Params{Spread: []stmt.Param{{Name: "ids"}}}}, true},
		{"with struct spread", &stmt.Query{Params: stmt.Params{StructSpread: []stmt.Struct{{Name: "User"}}}}, true},
		{"both spreads", &stmt.Query{Params: stmt.Params{
			Spread:       []stmt.Param{{Name: "ids"}},
			StructSpread: []stmt.Struct{{Name: "User"}},
		}}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, fn(tt.query))
		})
	}
}

func TestHasParams(t *testing.T) {
	fn := tmplfn["hasParams"].(func(*stmt.Query) bool)

	tests := []struct {
		name  string
		query *stmt.Query
		want  bool
	}{
		{"no params", &stmt.Query{}, false},
		{"scalar", &stmt.Query{Params: stmt.Params{Scalar: []stmt.Param{{Name: "id"}}}}, true},
		{"spread", &stmt.Query{Params: stmt.Params{Spread: []stmt.Param{{Name: "ids"}}}}, true},
		{"struct spread", &stmt.Query{Params: stmt.Params{StructSpread: []stmt.Struct{{Name: "User"}}}}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, fn(tt.query))
		})
	}
}

func TestValueToIdent(t *testing.T) {
	fn := tmplfn["valueToIdent"].(func(string) string)

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"simple", "active", "Active"},
		{"snake_case", "active_user", "ActiveUser"},
		{"with hyphen", "in-progress", "InProgress"}, // hyphen -> underscore -> camel
		{"with spaces", "hello world", "HelloWorld"}, // space -> underscore -> camel
		{"numbers", "status123", "Status123"},
		{"special chars", "status@#$", "Status"}, // special chars stripped then camelized
		{"trailing underscore", "active_", "Active"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, fn(tt.input))
		})
	}
}

func TestReturnsStruct(t *testing.T) {
	fn := tmplfn["returnsStruct"].(func(*stmt.Query) bool)

	tests := []struct {
		name  string
		query *stmt.Query
		want  bool
	}{
		{"exec mode", &stmt.Query{ExecMode: stmt.ExecModeExec, Returns: stmt.Struct{Params: []stmt.Param{{}, {}}}}, false},
		{"single return", &stmt.Query{ExecMode: stmt.ExecModeOne, Returns: stmt.Struct{Params: []stmt.Param{{}}}}, false},
		{"multiple returns", &stmt.Query{ExecMode: stmt.ExecModeOne, Returns: stmt.Struct{Params: []stmt.Param{{}, {}}}}, true},
		{"composite type", &stmt.Query{ExecMode: stmt.ExecModeOne, Returns: stmt.Struct{IsCompositeType: true, Params: []stmt.Param{{}}}}, true},
		{"many mode composite", &stmt.Query{ExecMode: stmt.ExecModeMany, Returns: stmt.Struct{IsCompositeType: true}}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, fn(tt.query))
		})
	}
}

func TestReturnsInlineStruct(t *testing.T) {
	fn := tmplfn["returnsInlineStruct"].(func(*stmt.Query) bool)

	tests := []struct {
		name  string
		query *stmt.Query
		want  bool
	}{
		{"exec mode", &stmt.Query{ExecMode: stmt.ExecModeExec, Returns: stmt.Struct{Params: []stmt.Param{{}, {}}}}, false},
		{"single return", &stmt.Query{ExecMode: stmt.ExecModeOne, Returns: stmt.Struct{Params: []stmt.Param{{}}}}, false},
		{"multiple returns inline", &stmt.Query{ExecMode: stmt.ExecModeOne, Returns: stmt.Struct{Params: []stmt.Param{{}, {}}}}, true},
		{"composite type", &stmt.Query{ExecMode: stmt.ExecModeOne, Returns: stmt.Struct{IsCompositeType: true, Params: []stmt.Param{{}, {}}}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, fn(tt.query))
		})
	}
}

func TestLowerFirstLetter(t *testing.T) {
	fn := tmplfn["lowerFirstLetter"].(func(string) string)

	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"empty", "", ""},
		{"uppercase", "UserID", "userID"},
		{"lowercase", "userId", "userId"},
		{"single char", "A", "a"},
		{"all caps", "ID", "iD"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, fn(tt.input))
		})
	}
}
