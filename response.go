package main

//Response - this is rest response
type Response struct {
	Success       bool           `json:"success,omitempty"`
	Message       string         `json:"message,omitempty"`
	Notifications []Notification `json:"notifications,omitempty"`
}
