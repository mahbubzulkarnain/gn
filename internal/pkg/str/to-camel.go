package str

import (
	"strings"
)

func ToCamel(s string) string {
	s = ToPascal(s)
	if len(s) == 0 {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}
