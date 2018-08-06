package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/say/{name}", say)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Welcome!")
}

func say(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	name := vars["name"]
	fmt.Fprintln(writer, "Hello, ", name)
}
