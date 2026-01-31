package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnakeToPascalCase(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"empty string", "", ""},
		{"single word", "user", "User"},
		{"two words", "user_name", "UserName"},
		{"three words", "first_middle_last", "FirstMiddleLast"},
		{"already capitalized", "User", "User"},
		{"uppercase word", "USER", "USER"},
		{"leading underscore", "_private", "Private"},
		{"trailing underscore", "name_", "Name"},
		{"multiple underscores", "a__b", "AB"},
		{"all underscores", "___", ""},

		// Common initialisms
		{"id initialism", "user_id", "UserID"},
		{"api initialism", "api_key", "APIKey"},
		{"http initialism", "http_url", "HTTPURL"},
		{"url initialism", "base_url", "BaseURL"},
		{"json initialism", "json_data", "JSONData"},
		{"sql initialism", "sql_query", "SQLQuery"},
		{"uuid initialism", "user_uuid", "UserUUID"},
		{"multiple initialisms", "http_api_url", "HTTPAPIURL"},
		{"initialism at start", "id_value", "IDValue"},
		{"initialism at end", "get_id", "GetID"},
		{"lowercase initialism match", "api", "API"},

		// Mixed cases
		{"mixed with numbers", "user_123", "User123"},
		{"number at start of word", "user_1name", "User1name"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, SnakeToPascalCase(tt.input))
		})
	}
}

func TestStructFieldNormalizer_Add(t *testing.T) {
	tests := []struct {
		name    string
		orig    string
		private bool
		want    string
		wantErr bool
	}{
		// Basic cases
		{"simple snake_case public", "user_name", false, "UserName", false},
		{"simple snake_case private", "user_name", true, "userName", false},
		{"single word public", "name", false, "Name", false},
		{"single word private", "name", true, "name", false},

		// Initialisms - private should lowercase entire leading initialism
		{"id public", "user_id", false, "UserID", false},
		{"id private", "user_id", true, "userID", false},
		{"api public", "api_key", false, "APIKey", false},
		{"api private", "api_key", true, "apiKey", false},
		{"dns_scoring private", "dns_scoring", true, "dnsScoring", false},
		{"ip_scoring private", "ip_scoring", true, "ipScoring", false},
		{"http_scoring private", "http_scoring", true, "httpScoring", false},
		{"tls_scoring private", "tls_scoring", true, "tlsScoring", false},
		{"ttl private", "ttl", true, "ttl", false},
		{"id alone private", "id", true, "id", false},
		{"url alone private", "url", true, "url", false},

		// Special characters (replaced with underscore then normalized)
		{"with hyphen", "user-name", false, "UserName", false},
		{"with dot", "user.name", false, "UserName", false},
		{"with space", "user name", false, "UserName", false},
		{"with special chars", "user@name#test", false, "UserNameTest", false},

		// Numbers at start
		{"starts with digit public", "123value", false, "X_123value", false},
		{"starts with digit private", "123value", true, "_123value", false},
		{"digit after underscore public", "_123", false, "X_123", false},
		{"digit after underscore private", "_123", true, "_123", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			snn := NewStructFieldNormalizer()
			got, err := snn.Add(tt.orig, tt.private)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestStructFieldNormalizer_Add_Duplicates(t *testing.T) {
	snn := NewStructFieldNormalizer()

	// First add should succeed
	got, err := snn.Add("user_name", false)
	assert.NoError(t, err)
	assert.Equal(t, "UserName", got)

	// Same original should fail (duplicate)
	got, err = snn.Add("user_name", false)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "duplicate parameter")
	assert.Equal(t, "UserName", got)

	// Different original that normalizes to same result should fail
	snn2 := NewStructFieldNormalizer()
	_, err = snn2.Add("user_name", false)
	assert.NoError(t, err)
	got, err = snn2.Add("user-name", false)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "duplicate parameter")
	assert.Contains(t, err.Error(), "user_name")
	assert.Contains(t, err.Error(), "user-name")
}

func TestStructFieldNormalizer_Add_EmptyResult(t *testing.T) {
	snn := NewStructFieldNormalizer()

	// Input that converts to empty string should error
	_, err := snn.Add("@#$%", false)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid identifier")
	assert.Contains(t, err.Error(), "converts to an empty string")
}

func TestStructFieldNormalizer_Add_MultipleFields(t *testing.T) {
	snn := NewStructFieldNormalizer()

	// Add multiple different fields
	fields := []struct {
		orig    string
		private bool
		want    string
	}{
		{"user_id", false, "UserID"},
		{"user_name", false, "UserName"},
		{"created_at", true, "createdAt"},
		{"api_key", false, "APIKey"},
	}

	for _, f := range fields {
		got, err := snn.Add(f.orig, f.private)
		assert.NoError(t, err)
		assert.Equal(t, f.want, got)
	}
}

func TestNewStructFieldNormalizer(t *testing.T) {
	snn := NewStructFieldNormalizer()
	assert.NotNil(t, snn)
	assert.NotNil(t, snn.keys)
	assert.Empty(t, snn.keys)
}

func TestToLowerCamelCase(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// Empty and simple cases
		{"empty string", "", ""},
		{"single lowercase", "a", "a"},
		{"single uppercase", "A", "a"},
		{"all lowercase", "user", "user"},

		// Normal PascalCase (single uppercase at start)
		{"normal word", "User", "user"},
		{"two words", "UserName", "userName"},
		{"three words", "FirstMiddleLast", "firstMiddleLast"},

		// Leading initialisms (multiple uppercase at start)
		{"ID alone", "ID", "id"},
		{"API alone", "API", "api"},
		{"URL alone", "URL", "url"},
		{"DNS alone", "DNS", "dns"},
		{"TTL alone", "TTL", "ttl"},

		// Initialism followed by word
		{"ID followed by word", "IDValue", "idValue"},
		{"API followed by word", "APIKey", "apiKey"},
		{"DNS followed by word", "DNSScoring", "dnsScoring"},
		{"IP followed by word", "IPScoring", "ipScoring"},
		{"HTTP followed by word", "HTTPScoring", "httpScoring"},
		{"TLS followed by word", "TLSScoring", "tlsScoring"},
		{"URL followed by word", "URLParser", "urlParser"},

		// Word followed by initialism (initialism not at start)
		{"word then ID", "UserID", "userID"},
		{"word then URL", "BaseURL", "baseURL"},
		{"word then API", "MyAPI", "myAPI"},

		// Multiple words with initialism in middle
		{"id in middle", "GetUserID", "getUserID"},

		// Edge cases
		{"all caps short", "AB", "ab"},
		{"all caps longer", "ABCD", "abcd"},
		{"mixed with numbers", "User123", "user123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, toLowerCamelCase(tt.input))
		})
	}
}
