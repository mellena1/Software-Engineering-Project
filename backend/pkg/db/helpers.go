package db

// StringPtr takes a string and returns a pointer to it
func StringPtr(myString string) *string {
	return &myString
}

// IntPtr takes an int and returns a pointer to it
func IntPtr(myInt int) *int {
	return &myInt
}

// Int64Ptr takes an int64 and returns a pointer to it
func Int64Ptr(myInt int64) *int64 {
	return &myInt
}
