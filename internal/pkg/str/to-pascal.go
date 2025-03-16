package str

import (
	"strings"
	"unicode"
)

func ToPascal(s string) string {
	if len(s) == 0 {
		return ""
	}

	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	var result strings.Builder
	for _, word := range words {
		if len(word) > 0 {
			word = strings.ToLower(word)
			result.WriteString(strings.ToUpper(word[:1]))
			if len(word) > 1 {
				result.WriteString(word[1:])
			}
		}
	}
	return result.String()
}
