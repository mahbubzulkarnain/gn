package str

import "github.com/mahbubzulkarnain/gn/internal/pkg/slug"

// ToSlug ...
func ToSlug(s string) string {
	return slug.Make(s)
}
