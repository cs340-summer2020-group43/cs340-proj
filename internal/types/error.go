package types


import (
	"cs340/internal"
	"strings"
	"net/url"
	"database/sql"
)


type Error struct {
	Id		int
	Desc	string
}


var errorColumnNames = []string{
	"id",
	"description",
}


func (e Error) ToSQLColumnsString() string {
	return internal.ToSQLColumnsString(errorColumnNames)
}


func (e Error) GetFormValues(form *url.Values) []interface{} {
	var values []interface{}
	for _, columnName := range errorColumnNames {
		values = append(values, form.Get(columnName))
	}
	return values
}


func (e Error) Insert(form *url.Values, db *sql.DB) error {
	beginningQuery := "insert into " + form.Get("table") + " "

	var builder strings.Builder

	builder.WriteString(beginningQuery)
	builder.WriteString(e.ToSQLColumnsString())
	builder.WriteString(" values(")

	if len(errorColumnNames) > 1 {
		for i:=0; i<len(errorColumnNames)-1; i++ {
			builder.WriteString("?, ")
		}
	}

	builder.WriteString("?);")

	_, err := db.Exec(
		builder.String(),
		e.GetFormValues(form)...,
	)

	return err
}


func (e Error) Update(form *url.Values, db *sql.DB) error {
	beginningQuery := "update " + form.Get("table") + " set "

	var builder strings.Builder

	builder.WriteString(beginningQuery)

	if len(cellColumnNames) > 1 {
		for i:=0; i<len(cellColumnNames)-1; i++ {
			builder.WriteString("`")
			builder.WriteString(cellColumnNames[i])
			builder.WriteString("`")
			builder.WriteString(" = (?), ")
		}
	}

	builder.WriteString("`")
	builder.WriteString(errorColumnNames[len(errorColumnNames)-1])
	builder.WriteString("`")
	builder.WriteString(" = (?) ")
	builder.WriteString("where `id` = (?);")

	var values []interface{} = e.GetFormValues(form)
	values = append(values, form.Get("id"))

	_, err := db.Exec(
		builder.String(),
		values...,
	)

	return err
}
