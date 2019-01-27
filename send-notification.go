package main

import (
	"log"
	"net/smtp"
	"strings"
)

// SendMail - sends mail
func SendMail(notification Notification) error {
	username := "*********"
	password := "*********"
	from := "no-reply@gmail.com"
	to := notification.To
	tostr := strings.Join(to, ",")
	subject := notification.Subject
	body := notification.Message
	msg := []byte("To:" + tostr + "\r\n" +
		"Subject:" + subject + "\r\n" +
		body + "\r\n")

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", username, password, "smtp.gmail.com"),
		from,
		to,
		msg)

	if err != nil {
		log.Fatal("An error occurred during sending mail.. ", err)
		return err
	}
	log.Print("Sent mail successfully")
	return nil
}
