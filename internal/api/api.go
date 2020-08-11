package api

import (
	"cs340/internal/types"
	"database/sql"
	"net/url"
)

func Insert(entry types.TableEntry, form *url.Values, db *sql.DB) error {
	return entry.Insert(form, db)
}

func Delete(form *url.Values, db *sql.DB) error {
	partialQuery := "delete from " + form.Get("table") + " where id=(?)"

	_, err := db.Exec(
		partialQuery,
		form.Get("id"),
	)

	return err
}

func Update(entry types.TableEntry, form *url.Values, db *sql.DB) error {
	return entry.Update(form, db)
}
