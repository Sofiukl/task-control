package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

/**dBurl - database URL
 */
const (
	dBurl                = "192.168.99.100:27017"
	notificationDatabase = "notificationdb"
)

// NewSession - its creates a new session
func _newSession() *mgo.Session {
	session, err := mgo.Dial(dBurl)
	if err != nil {
		log.Fatal("Fail to create session %s\n", err)
	}
	fmt.Println(session)
	return session
}

// GetNotificationDatabase - returns the notification database
func GetNotificationDatabase() *mgo.Database {
	Session := _newSession()
	return Session.DB(notificationDatabase)
}
