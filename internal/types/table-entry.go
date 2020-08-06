package types


import (
	"net/url"
	"database/sql"
)


type TableEntry interface {
	ToSQLColumnsString() string
	InsertIntoTable(*url.Values, *sql.DB) error
}
