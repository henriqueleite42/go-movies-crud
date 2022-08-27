package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go-movies-crud/src/ltypes"
	"go-movies-crud/src/routes"
)

func main() {
	r := mux.NewRouter()

	var moviesDatabase ltypes.MoviesDatabase

	moviesDatabase.NewMoviesDatabase()

	r.HandleFunc("/movies", routes.GetMovies(&moviesDatabase)).Methods("GET")
	r.HandleFunc("/movies/{id}", routes.GetMovie(&moviesDatabase)).Methods("GET")
	r.HandleFunc("/movies", routes.CreateMovie(&moviesDatabase)).Methods("POST")
	r.HandleFunc("/movies/{id}", routes.EditMovie(&moviesDatabase)).Methods("PUT")
	r.HandleFunc("/movies/{id}", routes.DeleteMovie(&moviesDatabase)).Methods("DELETE")

	fmt.Printf("Starting server at port 8080\n")

	log.Fatal(http.ListenAndServe(":8080", r))
}
