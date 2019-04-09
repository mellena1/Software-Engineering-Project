package mysql

import (
	"database/sql"
)

type rowScanner interface {
	Scan(dest ...interface{}) error
}

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

func NullStringToString(myString sql.NullString) *string {
	if myString.Valid {
		return &myString.String
	}
	return nil
}

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

func NullIntToInt(myInt sql.NullInt64) *int64 {
	if myInt.Valid {
		return &myInt.Int64
	}
	return nil
}
