package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Author struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type Course struct {
	Name   string  `json:"name"`
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

var courses []Course

func main() {
	fmt.Println("Hello World")

	//preparing some basic data
	courses = append(courses, Course{Name: "Jyothis", Id: 1, Title: "Atomic Habitats", Author: &Author{Name: "John", Id: 1}})
	courses = append(courses, Course{Name: "Arun", Id: 2, Title: "Atomic Habitats2", Author: &Author{Name: "Josan", Id: 2}})

	//initialize mux router
	r := mux.NewRouter()

	//define routers
	r.HandleFunc("/", ServeHome)
	r.HandleFunc("/allCourses", AllCourses)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hi Everyone!"))
}

func AllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(courses)
}
