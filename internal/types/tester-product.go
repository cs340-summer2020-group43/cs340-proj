package types

import (
	"cs340/internal"
	"database/sql"
	"net/url"
	"strings"
)

type TesterProduct struct {
	Tester	int
	Product	string
}

var testerProductColumnNames = []string{
	"tester",
	"product",
}

func (tp TesterProduct) ToSQLColumnsString() string {
	return internal.ToSQLColumnsString(testerColumnNames)
}

func (tp TesterProduct) GetFormValues(form *url.Values) []interface{} {
	var values []interface{}
	for _, columnName := range testerColumnNames {
		values = append(values, form.Get(columnName))
	}
	return values
}

func (tp TesterProduct) Insert(form *url.Values, db *sql.DB) error {
	beginningQuery := "insert into " + form.Get("table") + " "

	var builder strings.Builder

	builder.WriteString(beginningQuery)
	builder.WriteString(tp.ToSQLColumnsString())
	builder.WriteString(" values(")

	if len(testerColumnNames) > 1 {
		for i := 0; i < len(testerColumnNames)-1; i++ {
			builder.WriteString("?, ")
		}
	}

	builder.WriteString("?);")

	_, err := db.Exec(
		builder.String(),
		tp.GetFormValues(form)...,
	)

	return err
}

func (tp TesterProduct) Update(form *url.Values, db *sql.DB) error {
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
	builder.WriteString(testerColumnNames[len(testerColumnNames)-1])
	builder.WriteString("`")
	builder.WriteString(" = (?) ")
	builder.WriteString("where `id` = (?);")

	var values []interface{} = tp.GetFormValues(form)
	values = append(values, form.Get("id"))

	_, err := db.Exec(
		builder.String(),
		values...,
	)

	return err
}


func (tp TesterProduct) Delete(form *url.Values, db *sql.DB) error {
	var err error
	_, err = db.Exec(
		"delete from "+form.Get("table")+" where tester=(?) and product=(?);",
		form.Get("tester"),
		form.Get("product"),
	)
	return err
}
