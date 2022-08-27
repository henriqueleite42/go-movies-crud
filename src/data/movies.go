package data

import (
	"go-movies-crud/src/ltypes"
)

func GetMovies() []ltypes.Movie {
	var movies []ltypes.Movie

	movies = append(movies, ltypes.Movie{
		ID: "1",
		Isbn: "438227",
		Title: "Movie One",
		Director: &ltypes.Director{
			FirstName: "John",
			LastName: "Doe",
		},
	})
	movies = append(movies, ltypes.Movie{
		ID: "2",
		Isbn: "45455",
		Title: "Movie Two",
		Director: &ltypes.Director{
			FirstName: "Steve",
			LastName: "Smith",
		},
	})

	return movies
}