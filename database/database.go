package database

import (
	"database/sql"
)

type Database struct {
	path string
	db   *sql.DB

	artistsTable    *ArtistsTable
	componistsTable *ComponistsTable
}

func NewDatabase(path string) *Database {
	db := Database{path: path}

	db.artistsTable = NewArtistTable(&db)
	db.componistsTable = NewComponistTable(&db)

	return &db
}

// Open a database
func (db *Database) Open() error {
	var err error

	db.db, err = sql.Open("sqlite3", db.path)

	return err
}

// Init the database
// If the database is empty, the correct tabels are created
func (db *Database) Init() error {
	db.MustOpen()

	if err := db.artistsTable.Init(); err != nil { //nolint This way it is easy to add extra tables in the future without the need to verify older code
		return err
	}

	if err := db.componistsTable.Init(); err != nil {
		return err
	}

	return nil
}

func (db *Database) Close() error {
	if db.IsOpen() {
		return db.db.Close()
	}

	return nil
}

func (db *Database) Prepare() error {
	if err := db.Open(); err != nil {
		return err
	}

	defer func() {
		if err := db.Close(); err != nil {
			panic(err)
		}
	}()

	err := db.Init()
	return err
}

func (db *Database) IsOpen() bool {
	return db.db != nil
}

func (db *Database) MustOpen() {
	if !db.IsOpen() {
		panic("Database is not open!")
	}
}

func (db *Database) GetArtistsTable() *ArtistsTable {
	return db.artistsTable
}
