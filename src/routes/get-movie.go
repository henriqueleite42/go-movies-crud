package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"go-movies-crud/src/ltypes"
)

func GetMovie(movies []ltypes.Movie) func(w http.ResponseWriter, r *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)

		for _, item := range movies {
			if item.ID == params["id"] {
				json.NewEncoder(w).Encode(item)
				return
			}
		}
	}
}