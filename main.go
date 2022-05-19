package main

import (
	"fmt"
	"net/http"
	"log"
)

func main(){
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formhandle)
	http.HandleFunc("/hello", hellohandle)


	log.Println("Listening... on port `8080`")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func hellohandle(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path != "/hello"){
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path)
}

func formhandle(w http.ResponseWriter, r *http.Request) {
	if(r.URL.Path != "/form"){
		http.NotFound(w, r)
		return
	}
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, "Hello, %s!", r.FormValue("name"))
}