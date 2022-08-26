package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Student struct {
	ID         string      `json: "id"`
	name       string      `json: "name"`
	email      string      `json: "Email`
	Department *Department `json: "Department"`
}

type Department struct {
	name string `json: "Name"`
	year string `json: year`
}

var students []Student

func getStudentsList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, name := range students {
		if name.ID == params["ID"] {
			json.NewEncoder(w).Encode(name)
			return
		}
	}
}

func AddStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student Student
	_ = json.NewDecoder(r.Body).Decode(&student)
	student.ID = strconv.Itoa(rand.Intn(10000000))
	students = append(students, student)
	json.NewEncoder(w).Encode(student)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range students {
		if item.ID == params["ID"] {
			students = append(students[:index], students[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(students)
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range students {
		if item.ID == params["ID"] {
			students = append(students[:index], students[index+1:]...)
			var student Student
			_ = json.NewDecoder(r.Body).Decode(&student)
			student.ID = params["ID"]
			students = append(students, student)
			json.NewEncoder(w).Encode(student)
		}
	}
}
func main() {
	r := mux.NewRouter()

	students = append(students, Student{ID: "43", name: "Senthilnathan", email: "", Department: &Department{name: "AI&DS", year: "Second"}})
	students = append(students, Student{ID: "27", name: "", email: "", Department: &Department{name: "AI&DS", year: "Second"}})

	r.HandleFunc("/students", getStudentsList).Methods("GET")

	r.HandleFunc("/student/{id}", getStudent).Methods("GET")

	r.HandleFunc("/student/{id}", deleteStudent).Methods("POST")

	r.HandleFunc("/students", AddStudent).Methods("POST")

	r.HandleFunc("/student/{id}", updateStudent).Methods("POST")

	fmt.Printf("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
