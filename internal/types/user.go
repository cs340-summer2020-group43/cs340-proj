package types

import (
	"cs340/internal"
	"database/sql"
	"net/url"
	"strings"
)

type User struct {
	Id   int
	Name string
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

func (u User) Insert(form *url.Values, db *sql.DB) error {
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

func (u User) Update(form *url.Values, db *sql.DB) error {
	beginningQuery := "update " + form.Get("table") + " set "

	var builder strings.Builder

	builder.WriteString(beginningQuery)

	if len(cellColumnNames) > 1 {
		for i := 0; i < len(cellColumnNames)-1; i++ {
			builder.WriteString("`")
			builder.WriteString(cellColumnNames[i])
			builder.WriteString("`")
			builder.WriteString(" = (?), ")
		}
	}

	builder.WriteString("`")
	builder.WriteString(userColumnNames[len(userColumnNames)-1])
	builder.WriteString("`")
	builder.WriteString(" = (?) ")
	builder.WriteString("where `id` = (?);")

	var values []interface{} = u.GetFormValues(form)
	values = append(values, form.Get("id"))

	_, err := db.Exec(
		builder.String(),
		values...,
	)

	return err
}
