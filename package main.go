package main

import (

	"fmt"
	"net/http"
)

func handledFunc(w http.ResponseWriter, r *http.Request){

	fmt.Fprint(w, "<h1> Welcome to my site</h1>")



}


func main(){


	http.HandleFunc("/",handledFunc)
	http.ListenAndServe(":3200",nil)
}