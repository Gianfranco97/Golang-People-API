package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Person Basic people information
type Person struct {
	ID        int      `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address Basic address information
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

// Data Base simulation
var people []Person

// GetPeopleEndpoint Get all people
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// GetPersonEndpoint Get a person
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, _ := strconv.Atoi(params["id"])
	var res Person

	for _, item := range people {
		if item.ID == id {
			res = item
		}
	}

	json.NewEncoder(w).Encode(res)
}

// AddPersonEndpoint Add a new person
func AddPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	// params := mux.Vars(req)
	var person Person
	json.NewDecoder(req.Body).Decode(&person)
	person.ID = len(people) + 1

	people = append(people, person)
	json.NewEncoder(w).Encode(person)
}

func main() {
	router := mux.NewRouter()

	// Adding example data
	people = append(people, Person{ID: 1, FirstName: "Ryan", LastName: "Ray", Address: &Address{City: "Dubling", State: "California"}})
	people = append(people, Person{ID: 2, FirstName: "Maria", LastName: "Ray"})

	// Endpoints
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people", AddPersonEndpoint).Methods("POST")

	log.Fatal(http.ListenAndServe(":4000", router))
}
