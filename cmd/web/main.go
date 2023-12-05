package main

import (
	"fmt"
	"net/http"
	"github.com/krasnov23/web-app-on-golang/pkg/handlers"
)

 const portNumber = ":8080"

// main is the main app function
func main() {
	http.HandleFunc("/",handlers.Home)
	http.HandleFunc("/about",handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s",portNumber))
	http.ListenAndServe(portNumber,nil)
}

