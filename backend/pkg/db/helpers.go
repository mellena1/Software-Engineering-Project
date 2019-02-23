package db

// StringPtr takes a string and returns a pointer to it
func StringPtr(s string) *string {
	return &s
}

// IntPtr takes an int and returns a pointer to it
func IntPtr(i int) *int {
	return &i
}
