package types


import (
	"cs340/internal"
	"strings"
	"net/url"
	"database/sql"
)


type Cell struct {
	Id		int
	Desc	string
}


var cellColumnNames = []string{
	"id",
	"description",
}


func (c Cell) ToSQLColumnsString() string {
	return internal.ToSQLColumnsString(cellColumnNames)
}


func (c Cell) GetFormValues(form *url.Values) []interface{} {
	var values []interface{}
	for _, columnName := range cellColumnNames {
		values = append(values, form.Get(columnName))
	}
	return values
}


func (c Cell) Insert(form *url.Values, db *sql.DB) error {
	beginningQuery := "insert into " + form.Get("table") + " "

	var builder strings.Builder

	builder.WriteString(beginningQuery)
	builder.WriteString(c.ToSQLColumnsString())
	builder.WriteString(" values(")

	if len(cellColumnNames) > 1 {
		for i:=0; i<len(cellColumnNames)-1; i++ {
			builder.WriteString("?, ")
		}
	}

	builder.WriteString("?);")

	_, err := db.Exec(
		builder.String(),
		c.GetFormValues(form)...,
	)

	return err
}


func (c Cell) Update(form *url.Values, db *sql.DB) error {
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
	builder.WriteString(cellColumnNames[len(cellColumnNames)-1])
	builder.WriteString("`")
	builder.WriteString(" = (?) ")
	builder.WriteString("where `id` = (?);")

	var values []interface{} = c.GetFormValues(form)
	values = append(values, form.Get("id"))

	_, err := db.Exec(
		builder.String(),
		values...,
	)

	return err
}
