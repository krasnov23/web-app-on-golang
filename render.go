package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemple, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemple.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}