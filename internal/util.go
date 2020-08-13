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

var ManyToManyTableMap = map[string][]string{
	"Testers_Products": {"Testers", "Products"},
	"Testers_Users":    {"Testers", "Users"},
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

func StringInStrings(s string, strings []string) bool {
	for _, name := range strings {
		if s == name {
			return true
		}
	}
	return false
}

func IsTableName(s string) bool {
	return StringInStrings(s, TableNames)
}
