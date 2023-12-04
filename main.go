package main

import (
	"fmt"
	"html/template"
	"net/http"
)

 const portNumber = ":8080"

// Home page handler
func Home(w http.ResponseWriter,r *http.Request){
	renderTemplate(w,"home.page.tmpl")
}


// About is the about page handler
func About(w http.ResponseWriter,r *http.Request){
	renderTemplate(w,"about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemple,_ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemple.Execute(w,nil)
	if err != nil{
		fmt.Println("Error parsing template:",err)
		return
	}
}


// main is the main app function
func main() {
	http.HandleFunc("/",Home)
	http.HandleFunc("/about",About)

	fmt.Println(fmt.Sprintf("Starting application on port %s",portNumber))
	http.ListenAndServe(portNumber,nil)
}