package types


import (
	"fmt"
	"errors"
	"net/url"
	"database/sql"
)


type TableEntry interface {
	ToSQLColumnsString() string
	InsertIntoTable(*url.Values, *sql.DB) error
}


func TableEntryInit(tableName string) (TableEntry, error) {
	var tableEntry TableEntry
	var err error = nil

	switch tableName {
	case "Cells":
		tableEntry = Cell{}
	case "Errors":
		tableEntry = Error{}
	case "Products":
		tableEntry = Product{}
	case "Testers":
		tableEntry = Tester{}
	case "Tests":
		tableEntry = Test{}
	case "Users":
		tableEntry = User{}
	default:
		err = errors.New(fmt.Sprintf("Could not create TableEntry from \"%s\"", tableName))
	}

	return tableEntry, err
}
