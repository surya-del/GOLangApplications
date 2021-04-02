package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Student struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var stu []Student

func GetStuIdEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range stu {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Student{})
}
func GetStudentEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(stu)
}
func CreatestudentEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Student
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	stu = append(stu, person)
	json.NewEncoder(w).Encode(stu)
}
func DeleteStudentEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range stu {
		if item.ID == params["id"] {
			stu = append(stu[:index], stu[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(stu)
}
func main() {
	router := mux.NewRouter()
	stu = append(stu, Student{ID: "1", Firstname: "Nic", Lastname: "Raboy",
		Address: &Address{City: "Dublin", State: "CA"}})
	stu = append(stu, Student{ID: "2", Firstname: "Maria", Lastname: "Raboy"})

	router.HandleFunc("/student", GetStudentEndpoint).Methods("GET")
	router.HandleFunc("/student/{id}", GetStuIdEndpoint).Methods("GET")
	router.HandleFunc("/student/{id}", CreatestudentEndpoint).Methods("POST")
	router.HandleFunc("/student/{id}", DeleteStudentEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))
}
