package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	PORT = ":8000"
)

func main() {
	fmt.Println("Server is started at ", PORT)
	router := mux.NewRouter()
	ConfigureRouting(router)

	populateSomeData()
	// setting rest end points

	log.Fatal(http.ListenAndServe(PORT, router))

}
