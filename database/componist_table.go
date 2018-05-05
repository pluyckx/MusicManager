package database

const createComponistsTable = "CREATE TABLE IF NOT EXISTS componists (artist INTEGER PRIMARY KEY, `join` TEXT, UNIQUE (artist, `join`), FOREIGN KEY (artist) REFERENCES artists ( _rowid_ ) );"
const addComponist = "INSERT INTO componists (artist, join) VALUES($1, $2);"

type ComponistsTable struct {
	base baseTable
}

func NewComponistTable(db *Database) *ComponistsTable {
	return &ComponistsTable{base: baseTable{database: db}}
}

func (at *ComponistsTable) Init() error {
	at.base.database.MustOpen()

	_, err := at.base.database.db.Exec(createComponistsTable)

	return err
}

func (at *ComponistsTable) Add(artist uint, join string) error {
	at.base.database.MustOpen()

	_, err := at.base.database.db.Exec(addComponist, artist, join)

	return err
}
