package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"go-movies-crud/src/ltypes"
)

func GetMovie(moviesDatabase *ltypes.MoviesDatabase) func(w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)

		movie, err := moviesDatabase.FindOne(params["id"])

		if err != nil {
			http.Error(w, "Movie not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(movie)
	}
}