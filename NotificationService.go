package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

/**CollectionName = notification collection name
 */
const (
	CollectionName = "Notification"
)

/*CreateOne - create one notification
 */
func CreateOne(notification *Notification) error {
	Database := GetNotificationDatabase()
	return Database.C(CollectionName).Insert(notification)
}

/*UpdateOne - update one notification
 */
func UpdateOne(id string, status string) error {
	Database := GetNotificationDatabase()
	colQuerier := bson.M{"id": id}
	change := bson.M{"$set": bson.M{"status": status, "senttime": time.Now()}}
	return Database.C(CollectionName).Update(colQuerier, change)
}

// FindByID - Its find one document by ID in Notification collection
func FindByID(id string) (Notification, error) {
	var notification Notification
	Database := GetNotificationDatabase()
	err := Database.C(CollectionName).FindId(bson.ObjectIdHex(id)).One(&notification)
	if err != nil {
		return Notification{}, err
	}
	return notification, nil
}

/*FindAll - Its find all the documents in Notification collection
 */
func FindAll() ([]Notification, error) {
	var notifications []Notification
	Database := GetNotificationDatabase()
	err := Database.C(CollectionName).Find(nil).All(&notifications)
	if err != nil {
		return nil, err
	}
	return notifications, nil
}
