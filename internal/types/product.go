package types

import (
	"cs340/internal"
	"database/sql"
	"net/url"
	"strings"
)

type Product struct {
	Id            string
	Product_class string
	Kind          string
	Cell          int
}

var productColumnNames = []string{
	"id",
	"product_class",
	"type",
	"cell",
}

func (p Product) ToSQLColumnsString() string {
	return internal.ToSQLColumnsString(productColumnNames)
}

func (p Product) GetFormValues(form *url.Values) []interface{} {
	var values []interface{}
	for _, columnName := range productColumnNames {
		values = append(values, form.Get(columnName))
	}
	return values
}

func (p Product) Insert(form *url.Values, db *sql.DB) error {
	beginningQuery := "insert into " + form.Get("table") + " "

	var builder strings.Builder

	builder.WriteString(beginningQuery)
	builder.WriteString(p.ToSQLColumnsString())
	builder.WriteString(" values(")

	if len(productColumnNames) > 1 {
		for i := 0; i < len(productColumnNames)-1; i++ {
			builder.WriteString("?, ")
		}
	}

	builder.WriteString("?);")

	_, err := db.Exec(
		builder.String(),
		p.GetFormValues(form)...,
	)

	return err
}

func (p Product) Update(form *url.Values, db *sql.DB) error {
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
	builder.WriteString(productColumnNames[len(productColumnNames)-1])
	builder.WriteString("`")
	builder.WriteString(" = (?) ")
	builder.WriteString("where `id` = (?);")

	var values []interface{} = p.GetFormValues(form)
	values = append(values, form.Get("id"))

	_, err := db.Exec(
		builder.String(),
		values...,
	)

	return err
}
