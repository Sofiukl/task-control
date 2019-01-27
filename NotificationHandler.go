package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var notifications []Notification

//SendNotification -
// 1. create notification in the database with status : Not Send
// 2. send mail
// 3. update mail sent status to Sent in database
// 4. in case of failure update status to Failed
func SendNotification(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var notification Notification
	_ = json.NewDecoder(req.Body).Decode(&notification)
	notification.ID = params["id"]
	notification.Status = "Not Sent"
	notification.Senttime = time.Time{}
	fmt.Printf("Notification will be created.. %+v\n", notification)
	err := CreateOne(&notification)
	if err != nil {
		log.Fatal("Error occurred during write in the database..", err)
	}
	notificationErr := SendMail(notification)
	if notificationErr != nil {
		log.Fatal("Error occurred during write in the database..", notificationErr)
		response := Response{Success: false, Message: "Fail to sent notification", Notifications: []Notification{}}
		log.Printf("Response.. %+v\n", response)
		json.NewEncoder(res).Encode(response)
		return
	}
	//update Status in DB
	if err := UpdateOne(notification.ID, "Sent"); err != nil {
		//Update Status as Failed
		_ = UpdateOne(notification.ID, "Failed")
		response := Response{Success: false, Message: "Fail to sent notification", Notifications: []Notification{}}
		log.Printf("Response.. %+v\n", response)
		json.NewEncoder(res).Encode(response)
		return
	}
	notification.Status = "Sent"
	response := Response{Success: true, Message: "Successfully sent notification", Notifications: []Notification{notification}}
	log.Printf("Response.. %+v\n", response)
	json.NewEncoder(res).Encode(response)
}
