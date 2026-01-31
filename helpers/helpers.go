package helpers

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// commonInitialisms contains common programming initialisms that should be
// fully uppercased when converting from snake_case to PascalCase.
// Based on https://github.com/golang/go/wiki/CodeReviewComments#initialisms
var commonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
}

// SnakeToPascalCase converts a snake_case string to PascalCase.
// It handles common initialisms (e.g., "user_id" -> "UserID", "http_url" -> "HTTPURL").
func SnakeToPascalCase(s string) string {
	var result strings.Builder
	words := strings.Split(s, "_")

	for _, word := range words {
		if word == "" {
			continue
		}

		upper := strings.ToUpper(word)
		if commonInitialisms[upper] {
			result.WriteString(upper)
			continue
		}

		runes := []rune(word)
		runes[0] = unicode.ToUpper(runes[0])
		result.WriteString(string(runes))
	}

	return result.String()
}

// Matches all characters that can't be used in golang's identifiers
// https://golang.org/ref/spec#Identifiers
var identFix = regexp.MustCompile(`[^\pL\pN_]`)

type StructFieldNormalizer struct {
	keys map[string][]string
}

func NewStructFieldNormalizer() *StructFieldNormalizer {
	return &StructFieldNormalizer{keys: make(map[string][]string)}
}

// Add returns a normalized (snake->camel) orig or error, if the
// resulting normalized key was already added before.
//
// If private is true, keys are being normalized to camel case
// with the first letter being lower-cased
func (snn *StructFieldNormalizer) Add(orig string, private bool) (string, error) {
	normalized := identFix.ReplaceAllString(orig, "_")
	normalized = SnakeToPascalCase(normalized)
	if len(normalized) == 0 {
		return "", fmt.Errorf("invalid identifier '%v' (converts to an empty string)", orig)
	}

	r, _ := utf8.DecodeRuneInString(normalized)
	if unicode.IsDigit(r) {
		if private {
			normalized = "_" + normalized
		} else {
			normalized = "X_" + normalized
		}
	}

	if private {
		normalized = strings.ToLower(normalized[:1]) + normalized[1:]
	}

	snn.keys[normalized] = append(snn.keys[normalized], orig)
	if len(snn.keys[normalized]) > 1 {
		return normalized, fmt.Errorf("duplicate parameter %v (%v)", normalized,
			strings.Join(snn.keys[normalized], ", "))
	}

	return normalized, nil
}
