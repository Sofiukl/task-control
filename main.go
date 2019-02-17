package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	port = ":8000"
)

func main() {
	fmt.Println("GO Env... ", os.Getenv("ENVIRONMENT"))
	fmt.Println("Server is started at ", port)
	router := mux.NewRouter()
	ConfigureRouting(router)
	log.Fatal(http.ListenAndServe(port, router))
}
