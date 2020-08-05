package main


import (
	"cs340"
	cs340Template "cs340/internal/template"
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

			data := cs340Template.HomeData

			//data := []types.HomeDataEntry{
				//{Url: "cell-add.html", Desc: "Adds a Cell"},
				//{Url: "cell.html", Desc: "Shows Cells"},
				//{Url: "cell-update.html", Desc: "Updates a Cell"},
				//{Url: "error-add.html", Desc: "Adds an Error"},
				//{Url: "error.html", Desc: "Shows Errors"},
				//{Url: "error-update.html", Desc: "Updates and Error"},
				//{Url: "product-add.html", Desc: "Adds a Product"},
				//{Url: "product.html", Desc: "Shows Products"},
				//{Url: "product-update.html", Desc: "Updates a Product"},
				//{Url: "test-add.html", Desc: "Adds a Test"},
				//{Url: "tester-add.html", Desc: "Adds a Tester"},
				//{Url: "tester.html", Desc: "Shows Testers"},
				//{Url: "tester-product-add.html", Desc: "Adds a Tester-Product relationship"},
				//{Url: "tester-product.html", Desc: "Shows Tester-Product relationships"},
				//{Url: "tester-product-update.html", Desc: "Updates a Tester-Product relationship"},
				//{Url: "tester-update.html", Desc: "Updates a Tester"},
				//{Url: "test.html", Desc: "Shows Tests"},
				//{Url: "test-update.html", Desc: "Updates a Test"},
				//{Url: "user-add.html", Desc: "Adds a User"},
				//{Url: "user.html", Desc: "Shows Users"},
				//{Url: "user-update.html", Desc: "Updates Users"},
			//}

			t.Execute(writer, data)
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

		urlBasePath + "/add",
		func(writer http.ResponseWriter, req *http.Request) {

			tableNames := map[string]bool{
				"Cells": true,
				"Errors": true,
				"Products": true,
				"Testers": true,
				"Testers_Products": true,
				"Testers_Users": true,
				"Tests": true,
				"Users": true,
			}

			if req.Method != http.MethodPost {
				return
			}

			if !tableNames[req.FormValue("table")] {
				panic(1)
			}

			fmt.Println(req.Form)

			//beginningQuery := "insert into " + req.FormValue("table")

			//fmt.Println(beginningQuery)

			//db.Exec(
				//beginningQuery + " (`id`, `product_class`, `type`, `cell`) values(?, ?, ?, ?);",
				//req.FormValue("id"),
				//req.FormValue("product_class"),
				//req.FormValue("type"),
				//req.FormValue("cell"),
			//)
	})



	fmt.Printf("Server started on port %v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
