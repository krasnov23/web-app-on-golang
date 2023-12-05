package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// Выводит файл по названию и наследует в файл base файл из папки templates
	parsedTemple, _ := template.ParseFiles("./templates/" + tmpl,"./templates/base.layout.tmpl")
	err := parsedTemple.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}