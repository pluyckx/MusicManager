package app

import (
	"github.com/pluyckx/MusicManager/database"
)

var db = database.NewDatabase("music.sqlite3")

func GetDatabase() *database.Database {
	return db
}
