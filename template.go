package main

import (
	"html/template"
	"log"
	"os"
)

//Parse - parsing the template
func Parse(path string, config Notification) {
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Print(err)
		return
	}

	f, err := os.Create(path)
	if err != nil {
		log.Println("create file: ", err)
		return
	}
	err = t.Execute(f, config)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
	f.Close()
}
