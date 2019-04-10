package mysql

import (
	"database/sql"
)

// rowScanner interface for sql.Row and sql.Rows
type rowScanner interface {
	Scan(dest ...interface{}) error
}

// StringToNullString takes in a string pointer and returns a sql.NullString
func StringToNullString(myString *string) sql.NullString {
	if myString == nil {
		return sql.NullString{
			String: "",
			Valid:  false,
		}
	}
	return sql.NullString{
		String: *myString,
		Valid:  true,
	}
}

// NullStringToString takes in a sql.NullString and returns a string pointer
func NullStringToString(myString sql.NullString) *string {
	if myString.Valid {
		return &myString.String
	}
	return nil
}

// IntToNullInt takes in an int64 pointer and returns a sql.NullInt64
func IntToNullInt(myInt *int64) sql.NullInt64 {
	if myInt == nil {
		return sql.NullInt64{
			Int64: 0,
			Valid: false,
		}
	}
	return sql.NullInt64{
		Int64: *myInt,
		Valid: true,
	}
}

// NullIntToInt takes in a sql.NullInt64 and returns an int pointer
func NullIntToInt(myInt sql.NullInt64) *int64 {
	if myInt.Valid {
		return &myInt.Int64
	}
	return nil
}
