package main

import (
	"net/http"
	"fmt"
	"html/template"
)

func generateHTML(w http.ResponseWriter, data interface{}, fileNames ... string) {
	var files []string
	for _, file := range fileNames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}
