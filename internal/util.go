package internal


import (
	"strings"
)


var TableNames = []string{
	"Cells",
	"Errors",
	"Products",
	"Testers",
	"Testers_Products",
	"Testers_Users",
	"Tests",
	"Users",
}


func ToSQLColumnsString(columnNames []string) string {
	var builder strings.Builder

	builder.WriteString("(")

	for _, name := range columnNames {
		builder.WriteString("`" + name + "`, ")
	}

	s := builder.String()

	s = strings.TrimSuffix(s, ", ")

	s += ")"

	return s
}


func IsTableName(s string) bool {
	for _, name := range TableNames {
		if s == name {
			return true
		}
	}
	return false
}
