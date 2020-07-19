package main


import (
	"fmt"
	"html/template"
	"net/http"
)


const (
	port int = 3000
	urlBasePath = "/cs340"
	templatesBasePath = "../templates"
)


func main() {

	// The main page
	http.HandleFunc(
		urlBasePath + "/",
		func(writer http.ResponseWriter, req *http.Request) {
			data := make(map[string]string)
			data["Title"] = "Test Page"
			data["Body"] = "It works!"

			var t *template.Template
			t = template.Must(t.ParseFiles(templatesBasePath + "/base.tmpl"))
			t.Execute(writer, data)
	})


	fmt.Printf("Server started on port %v\n", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
