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

	// populateSomeData()

	//db
	// to := []string{"sofikul.mallick786@gmail.com"}
	// notification := Notification{"N1", to, "notification", "created"}
	// CreateOne(&notification)
	//var notifications []Notification
	//notifications, _ = FindAll()

	// find one
	// notification1, _ := FindByID("5c4c6465eb3fb320bec0f751")
	// log.Printf("Notification.. %+v\n", notification1)

	// send mail
	// SendMail(notification1)

	log.Fatal(http.ListenAndServe(PORT, router))

}
