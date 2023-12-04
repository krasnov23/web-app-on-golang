package main

import (
	"net/http"
)

// Home page handler
func Home(w http.ResponseWriter,r *http.Request){
	renderTemplate(w,"home.page.tmpl")
}


// About is the about page handler
func About(w http.ResponseWriter,r *http.Request){
	renderTemplate(w,"about.page.tmpl")
}