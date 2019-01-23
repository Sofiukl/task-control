package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var people []Person

func populateSomeData() {
	people = append(people, Person{"P1", "Sofikul Mallick", 28})
	people = append(people, Person{"P2", "Firoja Khatun", 23})
	people = append(people, Person{"P3", "Mehreen Sultana", 2})
	people = append(people, Person{"P4", "Mafuja Begum", 45})
}
func GetPeople(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(people)
}

func AddPerson(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	fmt.Printf("Person will be created.. %+v\n", person)
	people = append(people, person)
}

func GetPerson(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	ID := params["id"]
	for _, item := range people {
		if item.ID == ID {
			fmt.Printf("Person detail.. %+v\n", item)
			json.NewEncoder(res).Encode(item)
			return
		}
	}
	fmt.Printf("No person found with the specified ID %s\n", ID)
	json.NewEncoder(res).Encode(&Person{})
}

func DeletePerson(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	ID := params["id"]
	for index, person := range people {
		if person.ID == ID {
			fmt.Printf("Person deleted.. %+v\n", person)
			people = append(people[:index], people[index+1:]...)
			json.NewEncoder(res).Encode(person)
			return
		}
	}
	fmt.Printf("No person found with the specified ID %s\n", ID)
	json.NewEncoder(res).Encode(&Person{})
}
