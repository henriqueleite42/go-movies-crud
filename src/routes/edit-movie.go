package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"go-movies-crud/src/ltypes"
)

func EditMovie(moviesDatabase *ltypes.MoviesDatabase) func(w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)

		var movie ltypes.Movie
		_ = json.NewDecoder(r.Body).Decode(&movie)

		newMovie, err := moviesDatabase.Edit(params["id"], movie)

		if err != nil {
			http.Error(w, "Movie not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(newMovie)

	}
}