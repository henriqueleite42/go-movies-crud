package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"go-movies-crud/src/ltypes"
)

func DeleteMovie(moviesDatabase *ltypes.MoviesDatabase) func(w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)

		movies := moviesDatabase.Delete(params["id"])

		json.NewEncoder(w).Encode(movies)
	}
}