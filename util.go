package schemagen

import (
	"strings"
	"unicode"
)

func SnakeToCamel(s string) string {
	var b strings.Builder

	var upNext bool
	for _, r := range s {
		if r == '_' {
			upNext = true
			continue
		}
		if upNext {
			upNext = false
			b.WriteRune(unicode.ToUpper(r))
			continue
		}

		b.WriteRune(r)
	}

	return b.String()
}

func CamelToSnake(s string) string {
	var b strings.Builder

	for i, r := range s {
		if unicode.IsUpper(r) {
			if i != 0 {
				b.WriteRune('_')
			}
			b.WriteRune(unicode.ToLower(r))
			continue
		}
		b.WriteRune(r)
	}

	return b.String()
}
