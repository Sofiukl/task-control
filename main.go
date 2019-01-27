package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

const (
	port = ":8000"
)

func main() {
	fmt.Println("Server is started at ", port)
	router := mux.NewRouter()
	ConfigureRouting(router)
	//explore text template
	Parse("example.txt", Notification{ID: "N90", Subject: "Test subject", Message: "Test body", Status: "Active", Senttime: time.Time{}})
	f, _ := os.Open("example.txt")
	io.Copy(os.Stdout, f)
	f.Close()
	log.Fatal(http.ListenAndServe(port, router))
}
