package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
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
	r.HandleFunc("/allCourse", AllCourses).Methods("GET")
	r.HandleFunc("/createCourse", CreateCourse).Methods("POST")
	r.HandleFunc("/updateCourse/{id}", UpdateCourse).Methods("PUT")
	r.HandleFunc("/deleteCourse/{id}", DeleteCourse).Methods("DELETE")

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

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var course Course
	json.NewDecoder(r.Body).Decode(&course)
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	update_id, _ := strconv.Atoi(params["id"])
	for index, item := range courses {
		if item.Id == update_id {
			var updatedCourse Course
			_ = json.NewDecoder(r.Body).Decode(&updatedCourse)
			courses[index] = updatedCourse
			json.NewEncoder(w).Encode(courses[index])
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	update_id, _ := strconv.Atoi(params["id"])
	for index, item := range courses {
		if item.Id == update_id {
			courses = append(courses[:index], courses[index+1:]...)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
