package ptr

// Int returns a pointer to the input int.
func Int(i int) *int {
	return &i
}
