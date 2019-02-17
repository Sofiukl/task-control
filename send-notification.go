package main

import (
	"log"
	"net/smtp"
	"strings"
)

const (
	username = "sofikul.mallick@yopmail.com"
	password = ""
	from     = "no-reply@gmail.com"
)

// SendMail - sends mail
func SendMail(notification Notification) error {
	to := notification.To
	tostr := strings.Join(to, ",")
	subject := notification.Subject
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte("To:" + tostr + "\r\n" +
		"Subject:" + subject + "\r\n" +
		mime + "\r\n" +
		notification.Message + "\r\n")
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

// SendMailUsingTemplate - sends mail
func SendMailUsingTemplate(notification Notification) error {
	to := notification.To
	tostr := strings.Join(to, ",")
	subject := notification.Subject
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	bodyHTML, _ := ParseTemplate("mail_template.html", notification)
	msg := []byte("To:" + tostr + "\r\n" +
		"Subject:" + subject + "\r\n" +
		mime + "\r\n" +
		bodyHTML + "\r\n")
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
