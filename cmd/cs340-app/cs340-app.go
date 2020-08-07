package main


import (
	"cs340"
	"cs340/internal"
	"cs340/internal/api"
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

	defaultPageNames := []string{
		"cell-add.html",
		"cell-update.html",
		"error-add.html",
		"error-update.html",
		"product-add.html",
		"product-update.html",
		"test-add.html",
		"tester-add.html",
		"tester-product-add.html",
		"tester-product.html",
		"tester-product-update.html",
		"tester-update.html",
		"test-update.html",
		"user-add.html",
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


	for _, pageName := range defaultPageNames {
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

			tableEntry, err := types.TableEntryInit(req.FormValue("table"))
			if err != nil { log.Println(err) }

			err = api.InsertIntoTable(tableEntry, &req.Form, db)
			if err != nil { log.Println(err) }
	})


	http.HandleFunc(
		urlBasePath + "/cell.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var cells []types.Cell
			var rows *sql.Rows

			rows, err = db.Query("select * from Cells;")
			if err != nil {
				log.Println(err)
				rows.Close()
				return
			}
			defer rows.Close()

			for rows.Next() {
				cell := types.Cell{}

				err = rows.Scan(
					&cell.Id,
					&cell.Desc,
				)

				if err != nil {
					log.Println(err)
				}

				cells = append(cells, cell)
			}
			var t *template.Template
			t = template.Must(t.ParseFiles(cs340.TemplateBasePath + "/cell.tmpl"))
			t.Execute(writer, cells)
	})


	http.HandleFunc(
		urlBasePath + "/product.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var products []types.Product
			var rows *sql.Rows

			rows, err = db.Query("select * from Products;")
			if err != nil {
				log.Println(err)
				rows.Close()
				return
			}
			defer rows.Close()

			for rows.Next() {
				product := types.Product{}

				err = rows.Scan(
					&product.Id,
					&product.Product_class,
					&product.Kind,
					&product.Cell,
				)

				if err != nil {
					log.Println(err)
				}

				products = append(products, product)
			}
			var t *template.Template
			t = template.Must(t.ParseFiles(cs340.TemplateBasePath + "/product.tmpl"))
			t.Execute(writer, products)
	})


	http.HandleFunc(
		urlBasePath + "/error.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var errors []types.Error
			var rows *sql.Rows

			rows, err = db.Query("select * from Errors;")
			if err != nil {
				log.Println(err)
				rows.Close()
				return
			}
			defer rows.Close()

			for rows.Next() {
				e := types.Error{}

				err = rows.Scan(
					&e.Id,
					&e.Desc,
				)

				if err != nil {
					log.Println(err)
				}

				errors = append(errors, e)
			}
			var t *template.Template
			t = template.Must(t.ParseFiles(cs340.TemplateBasePath + "/error.tmpl"))
			t.Execute(writer, errors)
	})


	http.HandleFunc(
		urlBasePath + "/test.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var tests []types.Test
			var rows *sql.Rows

			rows, err = db.Query("select * from Tests;")
			if err != nil {
				log.Println(err)
				rows.Close()
				return
			}
			defer rows.Close()

			for rows.Next() {
				test := types.Test{}

				err = rows.Scan(
					&test.Id,
					&test.Product_id,
					&test.Serial_num,
					&test.Cell,
					&test.Start_time,
					&test.End_time,
					&test.Error_time,
					&test.Error_type,
					&test.Result_1,
					&test.Result_2,
					&test.Result_3,
					&test.Result_4,
					&test.Result_5,
					&test.Result_6,
					&test.Pass,
					&test.Tester,
					&test.User,
				)

				if err != nil {
					log.Println(err)
				}

				tests = append(tests, test)
			}
			var t *template.Template
			t = template.Must(t.ParseFiles(cs340.TemplateBasePath + "/test.tmpl"))
			t.Execute(writer, tests)
	})


	http.HandleFunc(
		urlBasePath + "/tester.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var testers []types.Tester
			var rows *sql.Rows

			rows, err = db.Query("select * from Testers;")
			if err != nil {
				log.Println(err)
				rows.Close()
				return
			}
			defer rows.Close()

			for rows.Next() {
				tester := types.Tester{}

				err = rows.Scan(
					&tester.Id,
					&tester.Desc,
					&tester.Cell,
				)

				if err != nil {
					log.Println(err)
				}

				testers = append(testers, tester)
			}
			var t *template.Template
			t = template.Must(t.ParseFiles(cs340.TemplateBasePath + "/tester.tmpl"))
			t.Execute(writer, testers)
	})


	http.HandleFunc(
		urlBasePath + "/user.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var users []types.User
			var rows *sql.Rows

			rows, err = db.Query("select * from Users;")
			if err != nil {
				log.Println(err)
				rows.Close()
				return
			}
			defer rows.Close()

			for rows.Next() {
				user := types.User{}

				err = rows.Scan(
					&user.Id,
					&user.Name,
				)

				if err != nil {
					log.Println(err)
				}

				users = append(users, user)
			}
			var t *template.Template
			t = template.Must(t.ParseFiles(cs340.TemplateBasePath + "/user.tmpl"))
			t.Execute(writer, users)
	})


	fmt.Printf("Server started on port %v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
