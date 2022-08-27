package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"go-movies-crud/src/ltypes"
)

func EditMovie(movies []ltypes.Movie) func(w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)

		for index, item := range movies {
			if item.ID == params["id"] {
				movies = append(movies[:index], movies[index+1:]...)

				var movie ltypes.Movie
				_ = json.NewDecoder(r.Body).Decode(&movie)
				movie.ID = params["id"]

				movies = append(movies, movie)

				json.NewEncoder(w).Encode(movie)
				return
			}
		}
	}
}