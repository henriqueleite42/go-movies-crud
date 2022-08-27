package ltypes

import (
	"errors"
)

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type MoviesDatabase struct {
	movies []Movie
}

func (moviesDatabase *MoviesDatabase) NewMoviesDatabase() {
	moviesDatabase.movies = append(moviesDatabase.movies, Movie{
		ID: "1",
		Isbn: "438227",
		Title: "Movie One",
		Director: &Director{
			FirstName: "John",
			LastName: "Doe",
		},
	})
	moviesDatabase.movies = append(moviesDatabase.movies, Movie{
		ID: "2",
		Isbn: "45455",
		Title: "Movie Two",
		Director: &Director{
			FirstName: "Steve",
			LastName: "Smith",
		},
	})
	moviesDatabase.movies = append(moviesDatabase.movies, Movie{
		ID: "3",
		Isbn: "666667",
		Title: "Movie Three",
		Director: &Director{
			FirstName: "José",
			LastName: "Bananas",
		},
	})
}

func (moviesDatabase *MoviesDatabase) Find() []Movie {
	return moviesDatabase.movies
}

func (moviesDatabase *MoviesDatabase) FindOne(id string) (Movie, error) {
	for _, item := range moviesDatabase.movies {
		if item.ID == id {
			return item, nil
		}
	}

	return Movie{}, errors.New("Movie not found")
}

func (moviesDatabase *MoviesDatabase) Add(movie Movie) []Movie {
	moviesDatabase.movies = append(moviesDatabase.movies, movie)

	return moviesDatabase.movies
}

func (moviesDatabase *MoviesDatabase) Delete(id string) []Movie {
	for index, item := range moviesDatabase.movies {
		if item.ID == id {
			moviesDatabase.movies = append(
				moviesDatabase.movies[:index],
				moviesDatabase.movies[index+1:]...
			)
			break
		}
	}

	return moviesDatabase.movies
}

func (moviesDatabase *MoviesDatabase) Edit(id string, data Movie) (Movie, error) {
	for index, item := range moviesDatabase.movies {
		if item.ID == id {
			moviesDatabase.movies = append(
				moviesDatabase.movies[:index],
				moviesDatabase.movies[index+1:]...
			)

			data.ID = id

			moviesDatabase.movies = append(moviesDatabase.movies, data)

			return data, nil
		}
	}

	return Movie{}, errors.New("Movie not found")
}