package database

const createArtistsTable = "CREATE TABLE IF NOT EXISTS artists (name TEXT UNIQUE);"
const addArtist = "INSERT INTO artists (name) VALUES($1);"

type ArtistsTable struct {
	base baseTable
}

func NewArtistTable(db *Database) *ArtistsTable {
	return &ArtistsTable{base: baseTable{database: db}}
}

func (at *ArtistsTable) Init() error {
	at.base.database.MustOpen()

	_, err := at.base.database.db.Exec(createArtistsTable)

	return err
}

func (at *ArtistsTable) Add(artist string) error {
	at.base.database.MustOpen()

	_, err := at.base.database.db.Exec(addArtist, artist)

	return err
}
