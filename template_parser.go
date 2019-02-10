package main

import (
	"bytes"
	"html/template"
	"log"
)

//ParseTemplate - parsing the template
func ParseTemplate(path string, data Notification) (string, error) {
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Print(err)
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		log.Println("Error to execute template..")
		return "", err
	}
	return buf.String(), nil
}
