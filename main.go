package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/pluyckx/MusicManager/app"
	"github.com/pluyckx/MusicManager/server"
)

func main() {
	db := app.GetDatabase()
	err := db.Prepare()
	if err != nil {
		panic(err)
	}

	err = server.ListenAndServe(12345)

	if err != nil {
		panic(err)
	}
}
