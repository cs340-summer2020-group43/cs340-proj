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
