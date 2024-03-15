package utils

import "database/sql"

// CreateNullString returns a sql.NullString with the given value, if not empty
func CreateNullString(value string) sql.NullString {
	if value != "" {
		return sql.NullString{String: value, Valid: true}
	}
	return sql.NullString{Valid: false}
}
