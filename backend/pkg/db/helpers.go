package db

func StringPtr(myString string) *string {
	return &myString
}

func IntPtr(myInt int) *int {
	return &myInt
}

func Int64Ptr(myInt int64) *int64 {
	return &myInt
}
