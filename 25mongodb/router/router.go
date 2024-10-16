package router

import (
	"github.com/gorilla/mux"
	"github.com/jyothisbenny/mongodb/controler"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controler.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controler.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controler.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controler.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controler.DeleteAllMovies).Methods("DELETE")
	return router
}
