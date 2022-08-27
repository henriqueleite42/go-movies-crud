package routes

import (
	"encoding/json"
	"net/http"

	"go-movies-crud/src/ltypes"
)

func GetMovies(movies []ltypes.Movie) func(w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(movies)
	}
}