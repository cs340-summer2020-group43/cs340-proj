package types


import (
	"cs340/internal"
	"strings"
	"net/url"
	"database/sql"
)


type Product struct {
	Id				string
	Product_class	string
	Kind			string
	Cell			int
}

var productColumnNames =[]string{
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
		for i:=0; i<len(productColumnNames)-1; i++ {
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
