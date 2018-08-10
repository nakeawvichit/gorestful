package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Welcome!!!")
}

func say(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	name := vars["name"]
	fmt.Fprintln(writer, "Hello, ", name)
}

func calljson(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(emps); err != nil {
		panic(err)
	}
}

func save(writer http.ResponseWriter, request *http.Request) {
	var emp Employee
	body, err := ioutil.ReadAll(io.LimitReader(request.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := request.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &emp); err != nil {
		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.WriteHeader(422)
		if err := json.NewEncoder(writer).Encode(err); err != nil {
			panic(err)
		}
	}
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(createEmployee(emp)); err != nil {
		panic(err)
	}
}
