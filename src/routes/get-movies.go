package routes

import (
	"encoding/json"
	"net/http"

	"go-movies-crud/src/ltypes"
)

func GetMovies(moviesDatabase *ltypes.MoviesDatabase) func(w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		movies := moviesDatabase.Find()

		json.NewEncoder(w).Encode(movies)
	}
}