package routes

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"go-movies-crud/src/ltypes"
)

func CreateMovie(moviesDatabase *ltypes.MoviesDatabase) func(w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var movie ltypes.Movie
		_ = json.NewDecoder(r.Body).Decode(&movie)

		movie.ID = strconv.Itoa(rand.Intn(100000000))

		moviesDatabase.Add(movie)

		json.NewEncoder(w).Encode(movie)
	}
}