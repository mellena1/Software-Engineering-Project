package mysql

import (
	"database/sql"
)

// rowScanner interface for sql.Row and sql.Rows
type rowScanner interface {
	Scan(dest ...interface{}) error
}

// StringToNullString takes in a string pointer and returns a sql.NullString
func StringToNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{
			String: "",
			Valid:  false,
		}
	}
	return sql.NullString{
		String: *s,
		Valid:  true,
	}
}

// NullStringToString takes in a sql.NullString and returns a string pointer
func NullStringToString(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}
	return nil
}

// IntToNullInt takes in an int64 pointer and returns a sql.NullInt64
func IntToNullInt(i *int64) sql.NullInt64 {
	if i == nil {
		return sql.NullInt64{
			Int64: 0,
			Valid: false,
		}
	}
	return sql.NullInt64{
		Int64: *i,
		Valid: true,
	}
}

// NullIntToInt takes in a sql.NullInt64 and returns an int pointer
func NullIntToInt(i sql.NullInt64) *int64 {
	if i.Valid {
		return &i.Int64
	}
	return nil
}
