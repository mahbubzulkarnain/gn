package ptr

// String returns a pointer to the input string.
func String(i string) *string {
	return &i
}
