package database

import (
	"github.com/mattn/go-sqlite3" //nolint Lint has problems with import
)

type baseTable struct {
	database *Database
}

func IsNotUniqueError(err error) bool {
	sqliteError, ok := err.(sqlite3.Error) //nolint sqlite3 has problems with importing

	return (ok && (sqliteError.ExtendedCode == sqlite3.ErrConstraintUnique)) //nolint sqlite3 has problems with importing
}
