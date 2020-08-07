package api


import (
	"cs340/internal/types"
	"net/url"
	"database/sql"
)


func InsertIntoTable(entry types.TableEntry,  form *url.Values, db *sql.DB) error {
	return entry.InsertIntoTable(form, db)
}


func DeleteFromTable(form *url.Values, db *sql.DB) error {
	partialQuery := "delete from " + form.Get("table") + " where id=(?)"

	_, err := db.Exec(
		partialQuery,
		form.Get("id"),
	)

	return err
}
