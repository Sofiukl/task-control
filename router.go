package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

type handleRequest func(http.ResponseWriter, *http.Request)

//ConfigureRouting - configures all the client routes
func ConfigureRouting(route *mux.Router) {
	router = route
	// AddARoute("GET", "/people", GetPeople)
	// AddARoute("GET", "/people/{id}", GetPerson)
	AddARoute("POST", "/notification/{id}", SendNotification)
	// AddARoute("DELETE", "/people/{id}", DeletePerson)
}

//AddARoute - function for attaching route handlers
func AddARoute(method string, path string, fn handleRequest) {
	router.HandleFunc(path, fn).Methods(method)
}
