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


func (c Cell) InsertIntoTable(form *url.Values, db *sql.DB) error {
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
