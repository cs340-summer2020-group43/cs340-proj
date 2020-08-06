package types


import (
	"cs340/internal"
	"strings"
	"net/url"
	"database/sql"
)


type User struct {
	Id		int
	Name	string
}


var userColumnNames = []string{
	"id",
	"name",
}


func (u User) ToSQLColumnsString() string {
	return internal.ToSQLColumnsString(userColumnNames)
}


func (u User) GetFormValues(form *url.Values) []interface{} {
	var values []interface{}
	for _, columnName := range userColumnNames {
		values = append(values, form.Get(columnName))
	}
	return values
}


func (u User) InsertIntoTable(form *url.Values, db *sql.DB) error {
	beginningQuery := "insert into " + form.Get("table") + " "

	var builder strings.Builder

	builder.WriteString(beginningQuery)
	builder.WriteString("(`name`)")
	builder.WriteString(" values(")

	builder.WriteString("?);")

	_, err := db.Exec(
		builder.String(),
		form.Get("name"),
	)

	return err
}
