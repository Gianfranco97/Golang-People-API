package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

// GetPersonEndpoint Get a person
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Fprintf(w, "Insert a valid ID")
		return
	}

	var res Person

	for _, item := range people {
		if item.ID == id {
			res = item
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(res)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "No items were found with that ID")
}

// AddPersonEndpoint Add a new person
func AddPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	var person Person
	json.NewDecoder(req.Body).Decode(&person)
	reqBody, err := ioutil.ReadAll(req.Body)

	if err != nil {
		fmt.Fprintf(w, "Insert a valid data")
		return
	}
	json.Unmarshal(reqBody, &person)

	person.ID = len(people) + 1

	people = append(people, person)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}

// DeletePersonEndpoint Delete a person
func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Fprintf(w, "Insert a valid ID")
		return
	}

	for index, item := range people {
		if item.ID == id {
			people = append(people[:index], people[index+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(people)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "No items were found with that ID")

}

func main() {
	router := mux.NewRouter()

	// Adding example data
	people = append(people, Person{ID: 1, FirstName: "Ryan", LastName: "Ray", Address: &Address{City: "Dubling", State: "California"}})
	people = append(people, Person{ID: 2, FirstName: "Maria", LastName: "Ray"})

	// Endpoints
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people", AddPersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", router))
}
