package main


import (
	"cs340"
	"cs340/internal"
	cs340Template "cs340/internal/template"
	"cs340/internal/types"
    "os"
	"fmt"
    "strconv"
	"html/template"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


const (
	urlBasePath = "/cs340"
)


// Usage
// -----
// app [PORT] [DB] [USER] [PASS]
// 
// Args
// ----
// PORT: the port to listen on
// DB: the name of the database to connect to
// USER, PASS: login info for the database
func main() {

    port, err := strconv.Atoi(os.Args[1])
    dbLocation := os.Args[3] + ":" + os.Args[4] + "@(localhost)/" + os.Args[2]

	db, err := sql.Open("mysql", dbLocation)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}


	var fileServer = http.FileServer(http.Dir("web/static/"))

	pageNames := []string{
		"cell-add.html",
		"cell.html",
		"cell-update.html",
		"error-add.html",
		"error.html",
		"error-update.html",
		"product-add.html",
		"product.html",
		"product-update.html",
		"test-add.html",
		"tester-add.html",
		"tester.html",
		"tester-product-add.html",
		"tester-product.html",
		"tester-product-update.html",
		"tester-update.html",
		"test.html",
		"test-update.html",
		"user-add.html",
		"user.html",
		"user-update.html",
	}


	http.Handle(urlBasePath + "/static/", http.StripPrefix(urlBasePath + "/static/", fileServer))


	http.HandleFunc(
		urlBasePath + "/",
		func(writer http.ResponseWriter, req *http.Request) {
			var t *template.Template
			t = template.Must(t.ParseFiles(cs340.TemplateBasePath + "/home.tmpl"))
			t.Execute(writer, cs340Template.HomeData)
	})


	for _, pageName := range pageNames {
		pagePath := "/" + pageName
		http.HandleFunc(
			urlBasePath + pagePath,
			func(writer http.ResponseWriter, req *http.Request) {
				var t *template.Template
				t = template.Must(t.ParseFiles(cs340.HtmlBasePath + pagePath))
				t.Execute(writer, "")
		})
	}


	http.HandleFunc(

		urlBasePath + "/insert",
		func(writer http.ResponseWriter, req *http.Request) {

			if req.Method != http.MethodPost {
				return
			}

			if !internal.IsTableName(req.FormValue("table")) {
				return
			}

			var tableEntry types.TableEntry

			switch req.FormValue("table") {
			case "Cells":
				tableEntry = types.Cell{}
			case "Errors":
				tableEntry = types.Error{}
			case "Products":
				tableEntry = types.Product{}
			case "Testers":
				tableEntry = types.Tester{}
			case "Tests":
				tableEntry = types.Test{}
			case "Users":
				tableEntry = types.User{}
			default:
				return
			}

			fmt.Println(tableEntry.ToSQLColumnsString())

			err := tableEntry.InsertIntoTable(&req.Form, db)

			if err != nil {
				log.Println(err)
			}
	})



	fmt.Printf("Server started on port %v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
