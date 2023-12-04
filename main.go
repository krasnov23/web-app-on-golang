package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter,r *http.Request){

}


func main() {
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		
		// Вывод hello world на экран
		n,err := fmt.Fprintf(w,"Hello world")
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written: %d",n))
	})

	http.ListenAndServe(":8080",nil)
}