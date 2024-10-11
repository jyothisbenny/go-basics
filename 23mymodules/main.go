package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello world")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	log.Fatal(http.ListenAndServe(":8080", r))

}

func greeter() {
	fmt.Println("greetings world!")

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello jyothis"))
}
