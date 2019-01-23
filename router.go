package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

type handleRequest func(http.ResponseWriter, *http.Request)

func ConfigureRouting(route *mux.Router) {
	router = route
	AddARoute("GET", "/people", GetPeople)
	AddARoute("GET", "/people/{id}", GetPerson)
	AddARoute("POST", "/people/{id}", AddPerson)
	AddARoute("DELETE", "/people/{id}", DeletePerson)

}

func AddARoute(method string, path string, fn handleRequest) {
	router.HandleFunc(path, fn).Methods(method)
}
