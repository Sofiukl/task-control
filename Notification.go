package main

import (
	"time"
)

/*Notification structure
* for primarily Email
 */
type Notification struct {
	ID       string
	To       []string
	Subject  string
	Message  string
	Status   string
	Senttime time.Time
}
