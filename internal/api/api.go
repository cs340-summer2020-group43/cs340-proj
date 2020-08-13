package api

import (
	"cs340/internal"
	"cs340/internal/types"
	"database/sql"
	"net/url"
	"strings"
	"unicode"
)

func Insert(entry types.TableEntry, form *url.Values, db *sql.DB) error {
	return entry.Insert(form, db)
}

func Delete(form *url.Values, db *sql.DB) error {

	if it := getIntersectionTable(form.Get("table")); it != "" {
		deleteManyToManyParticipation(it, form, db)
	}

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

func getIntersectionTable(s string) string {
	// it = intersection_table
	// pt = participating_table
	for it, pts := range internal.ManyToManyTableMap {
		for _, pt := range pts {
			if s == pt {
				return it
			}
		}
	}
	return ""
}

func deleteManyToManyParticipation(
	intersection_table string,
	form *url.Values,
	db *sql.DB,
) error {

	column_name := tableNameToParticipatoryColumnName(form.Get("table"))

	partialQuery := "delete from " + intersection_table +
		" where " + column_name + "=(?)"

	_, err := db.Exec(
		partialQuery,
		form.Get("id"),
	)

	return err
}

func tableNameToParticipatoryColumnName(n string) string {
	var column_name_b strings.Builder

	column_name_b.WriteByte(byte(unicode.ToLower(rune(n[0]))))

	for i := 1; i < len(n)-1; i++ {
		column_name_b.WriteByte(n[i])
	}

	return column_name_b.String()
}
