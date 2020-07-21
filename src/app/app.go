package main


import (
	"fmt"
	"html/template"
	"net/http"
)


const (
	templateBasePath = "./src/templates"
	htmlBasePath = "./src/static/html"
	port int = 4000
	urlBasePath = "/cs340"
)


func main() {

	http.HandleFunc(
		urlBasePath + "/",
		func(writer http.ResponseWriter, req *http.Request) {
			var t *template.Template
			t = template.Must(t.ParseFiles(templateBasePath + "/home.tmpl"))

			type HomeDataEntry struct {
				Url string
				Desc string
			}

			data := []HomeDataEntry{
				{Url: "cell-add.html", Desc: "Adds a Cell"},
				{Url: "cell.html", Desc: "Shows Cells"},
				{Url: "cell-update.html", Desc: "Updates a Cell"},
				{Url: "error-add.html", Desc: "Adds an Error"},
				{Url: "error.html", Desc: "Shows Erros"},
				{Url: "error-update.html", Desc: "Updates and Error"},
				{Url: "product-add.html", Desc: "Adds a Product"},
				{Url: "product.html", Desc: "Shows Products"},
				{Url: "product-update.html", Desc: "Updates a Product"},
				{Url: "test-add.html", Desc: "Adds a Test"},
				{Url: "tester-add.html", Desc: "Adds a Tester"},
				{Url: "tester.html", Desc: "Shows Testers"},
				{Url: "tester-product-add.html", Desc: "Adds a Tester-Product relationship"},
				{Url: "tester-product.html", Desc: "Shows Tester-Product relationships"},
				{Url: "tester-product-update.html", Desc: "Updates a Tester-Product relationship"},
				{Url: "tester-update.html", Desc: "Updates a Tester"},
				{Url: "test.html", Desc: "Shows Tests"},
				{Url: "test-update.html", Desc: "Updates a Test"},
				{Url: "user-add.html", Desc: "Adds a User"},
				{Url: "user.html", Desc: "Shows Users"},
				{Url: "user-update.html", Desc: "Updates Users"},
			}

			t.Execute(writer, data)
	})

	http.HandleFunc(
		urlBasePath + "/cell-add.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var t *template.Template
			t = template.Must(t.ParseFiles(htmlBasePath + "/cell-add.html"))
			t.Execute(writer, "")
	})

	http.HandleFunc(
		urlBasePath + "/cell.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var t *template.Template
			t = template.Must(t.ParseFiles(htmlBasePath + "/cell.html"))
			t.Execute(writer, "")
	})

	http.HandleFunc(
		urlBasePath + "/cell-update.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var t *template.Template
			t = template.Must(t.ParseFiles(htmlBasePath + "/cell-update.html"))
			t.Execute(writer, "")
	})

	http.HandleFunc(
		urlBasePath + "/error-add.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var t *template.Template
			t = template.Must(t.ParseFiles(htmlBasePath + "/error-add.html"))
			t.Execute(writer, "")
	})

	http.HandleFunc(
		urlBasePath + "/error.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var t *template.Template
			t = template.Must(t.ParseFiles(htmlBasePath + "/error.html"))
			t.Execute(writer, "")
	})

	http.HandleFunc(
		urlBasePath + "/error-update.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var t *template.Template
			t = template.Must(t.ParseFiles(htmlBasePath + "/error-update.html"))
			t.Execute(writer, "")
	})

	http.HandleFunc(
		urlBasePath + "/product-add.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var t *template.Template
			t = template.Must(t.ParseFiles(htmlBasePath + "/product-add.html"))
			t.Execute(writer, "")
	})

	http.HandleFunc(
		urlBasePath + "/product.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var t *template.Template
			t = template.Must(t.ParseFiles(htmlBasePath + "/product.html"))
			t.Execute(writer, "")
	})

	http.HandleFunc(
		urlBasePath + "/product-update.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var t *template.Template
			t = template.Must(t.ParseFiles(htmlBasePath + "/product-update.html"))
			t.Execute(writer, "")
	})

	http.HandleFunc(
		urlBasePath + "/test-add.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var t *template.Template
			t = template.Must(t.ParseFiles(htmlBasePath + "/test-add.html"))
			t.Execute(writer, "")
	})

	http.HandleFunc(
		urlBasePath + "/tester-add.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var t *template.Template
			t = template.Must(t.ParseFiles(htmlBasePath + "/tester-add.html"))
			t.Execute(writer, "")
	})

	http.HandleFunc(
		urlBasePath + "/tester.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var t *template.Template
			t = template.Must(t.ParseFiles(htmlBasePath + "/tester.html"))
			t.Execute(writer, "")
	})

	http.HandleFunc(
		urlBasePath + "/tester-product-add.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var t *template.Template
			t = template.Must(t.ParseFiles(htmlBasePath + "/tester-product-add.html"))
			t.Execute(writer, "")
	})

	http.HandleFunc(
		urlBasePath + "/tester-product.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var t *template.Template
			t = template.Must(t.ParseFiles(htmlBasePath + "/tester-product.html"))
			t.Execute(writer, "")
	})

	http.HandleFunc(
		urlBasePath + "/tester-product-update.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var t *template.Template
			t = template.Must(t.ParseFiles(htmlBasePath + "/tester-product-update.html"))
			t.Execute(writer, "")
	})

	http.HandleFunc(
		urlBasePath + "/tester-update.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var t *template.Template
			t = template.Must(t.ParseFiles(htmlBasePath + "/tester-update.html"))
			t.Execute(writer, "")
	})

	http.HandleFunc(
		urlBasePath + "/test.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var t *template.Template
			t = template.Must(t.ParseFiles(htmlBasePath + "/test.html"))
			t.Execute(writer, "")
	})

	http.HandleFunc(
		urlBasePath + "/test-update.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var t *template.Template
			t = template.Must(t.ParseFiles(htmlBasePath + "/test-update.html"))
			t.Execute(writer, "")
	})

	http.HandleFunc(
		urlBasePath + "/user-add.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var t *template.Template
			t = template.Must(t.ParseFiles(htmlBasePath + "/user-add.html"))
			t.Execute(writer, "")
	})

	http.HandleFunc(
		urlBasePath + "/user.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var t *template.Template
			t = template.Must(t.ParseFiles(htmlBasePath + "/user.html"))
			t.Execute(writer, "")
	})

	http.HandleFunc(
		urlBasePath + "/user-update.html",
		func(writer http.ResponseWriter, req *http.Request) {
			var t *template.Template
			t = template.Must(t.ParseFiles(htmlBasePath + "/user-update.html"))
			t.Execute(writer, "")
	})



	fmt.Printf("Server started on port %v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
