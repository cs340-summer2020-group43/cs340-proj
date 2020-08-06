package types


import (
	"cs340/internal"
	"strings"
	"net/url"
	"database/sql"
)


type Tester struct {
	Id		int
	Desc	string
	Cell	int
}


var testerColumnNames = []string{
	"id",
	"description",
	"cell",
}


func (t Tester) ToSQLColumnsString() string {
	return internal.ToSQLColumnsString(testerColumnNames)
}


func (t Tester) InsertIntoTable(form *url.Values, db *sql.DB) error {
	beginningQuery := "insert into " + form.Get("table") + " "

	var builder strings.Builder

	builder.WriteString(beginningQuery)
	builder.WriteString(t.ToSQLColumnsString())
	builder.WriteString(" values(")

	if len(testerColumnNames) > 1 {
		for i:=0; i<len(testerColumnNames)-1; i++ {
			builder.WriteString("?, ")
		}
	}

	builder.WriteString("?);")

	_, err := db.Exec(
		builder.String(),
		t.GetFormValues(form)...,
	)

	return err
}


func (t Tester) GetFormValues(form *url.Values) []interface{} {
	var values []interface{}
	for _, columnName := range testerColumnNames {
		values = append(values, form.Get(columnName))
	}
	return values
}
