package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go-movies-crud/src/data"
	"go-movies-crud/src/routes"
)

func main() {
	r := mux.NewRouter()

	movies := data.GetMovies()

	r.HandleFunc("/movies", routes.GetMovies(movies)).Methods("GET")
	r.HandleFunc("/movies/{id}", routes.GetMovie(movies)).Methods("GET")
	r.HandleFunc("/movies", routes.CreateMovie(movies)).Methods("POST")
	r.HandleFunc("/movies/{id}", routes.EditMovie(movies)).Methods("PUT")
	r.HandleFunc("/movies/{id}", routes.DeleteMovie(movies)).Methods("DELETE")

	fmt.Printf("Starting server at port 8080\n")

	log.Fatal(http.ListenAndServe(":8080", r))
}
