package helpers

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/serenize/snaker"
)

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
	normalized = snaker.SnakeToCamel(normalized)
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
