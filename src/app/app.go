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
				{Url: "error.html", Desc: "Shows Errors"},
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

	for _, pageName := range pageNames {
		pagePath := "/" + pageName
		http.HandleFunc(
			urlBasePath + pagePath,
			func(writer http.ResponseWriter, req *http.Request) {
				var t *template.Template
				t = template.Must(t.ParseFiles(htmlBasePath + pagePath))
				t.Execute(writer, "")
		})
	}



	fmt.Printf("Server started on port %v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
