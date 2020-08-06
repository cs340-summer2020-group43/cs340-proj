package types


import (
	"cs340/internal"
	"strings"
	"net/url"
	"database/sql"
)


type Test struct {
	Id			int
	Product_id	string
	Serial_num	string
	Cell		int
	Start_time	string
	End_time	string
	Error_time	string
	Error_type	int
	Result_1	float64
	Result_2	float64
	Result_3	float64
	Result_4	float64
	Result_5	float64
	Result_6	float64
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


func (t Test) InsertIntoTable(form *url.Values, db *sql.DB) error {
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
