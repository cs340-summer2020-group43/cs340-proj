package types


import (
	"cs340/internal"
	"strings"
	"net/url"
	"database/sql"
	"time"
)


type Test struct {
	Id			int
	Product_id	string
	Serial_num	string
	Cell		int
	Start_time	time.Time
	End_time	time.Time
	Error_time	sql.NullTime
	Error_type	sql.NullInt32
	Result_1	sql.NullFloat64
	Result_2	sql.NullFloat64
	Result_3	sql.NullFloat64
	Result_4	sql.NullFloat64
	Result_5	sql.NullFloat64
	Result_6	sql.NullFloat64
	Pass		int
	Tester		int
	User		int
}


var testColumnNames = []string{
	"id",
	"product_id",
	"serial_num",
	"cell",
	"start_time",
	"end_time",
	"error_time",
	"error_type",
	"result_1",
	"result_2",
	"result_3",
	"result_4",
	"result_5",
	"result_6",
	"pass",
	"tester",
	"user",
}


func (t Test) ToSQLColumnsString() string {
	return internal.ToSQLColumnsString(testColumnNames)
}


func (t Test) GetFormValues(form *url.Values) []interface{} {
	var values []interface{}
	for _, columnName := range testColumnNames {
		values = append(values, form.Get(columnName))
	}
	return values
}


func (t Test) Insert(form *url.Values, db *sql.DB) error {
	beginningQuery := "insert into " + form.Get("table") + " "

	var builder strings.Builder

	builder.WriteString(beginningQuery)
	builder.WriteString(t.ToSQLColumnsString())
	builder.WriteString(" values(")

	if len(testColumnNames) > 1 {
		for i:=0; i<len(testColumnNames)-1; i++ {
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


func (t Test) Update(form *url.Values, db *sql.DB) error {
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
	builder.WriteString(testColumnNames[len(testColumnNames)-1])
	builder.WriteString("`")
	builder.WriteString(" = (?) ")
	builder.WriteString("where `id` = (?);")

	var values []interface{} = t.GetFormValues(form)
	values = append(values, form.Get("id"))

	_, err := db.Exec(
		builder.String(),
		values...,
	)

	return err
}
