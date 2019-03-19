package mysql

import (
	"database/sql"
	"testing"

	"github.com/mellena1/Software-Engineering-Project/backend/pkg/db"

	"github.com/stretchr/testify/assert"
)

func TestStringToNullString(t *testing.T) {
	val := "hello"
	expected := sql.NullString{
		String: val,
		Valid:  true,
	}
	actual := stringToNullString(&val)
	assert.Equal(t, expected, actual)

	expected = sql.NullString{
		String: "",
		Valid:  false,
	}
	actual = stringToNullString(nil)
	assert.Equal(t, expected, actual)
}

func TestNullStringToString(t *testing.T) {
	val := sql.NullString{
		String: "hello",
		Valid:  true,
	}
	expected := db.StringPtr("hello")
	actual := nullStringToString(val)
	assert.Equal(t, expected, actual)

	val = sql.NullString{
		String: "",
		Valid:  false,
	}
	expected = nil
	actual = nullStringToString(val)
	assert.Equal(t, expected, actual)
}

func TestIntToNullInt(t *testing.T) {
	val := int64(2)
	expected := sql.NullInt64{
		Int64: val,
		Valid: true,
	}
	actual := intToNullInt(&val)
	assert.Equal(t, expected, actual)

	expected = sql.NullInt64{
		Int64: 0,
		Valid: false,
	}
	actual = intToNullInt(nil)
	assert.Equal(t, expected, actual)
}

func TestNullIntToInt(t *testing.T) {
	val := sql.NullInt64{
		Int64: 15,
		Valid: true,
	}
	expected := db.Int64Ptr(15)
	actual := nullIntToInt(val)
	assert.Equal(t, expected, actual)

	val = sql.NullInt64{
		Int64: 0,
		Valid: false,
	}
	expected = nil
	actual = nullIntToInt(val)
	assert.Equal(t, expected, actual)
}
