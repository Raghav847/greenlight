package data

import (
	"time"

	"github.com/Raghav847/greenlight/internal/validator"
)

type Movie struct {
	ID        int64     `json:"id"`                //unique int ID for the movie
	CreatedAt time.Time `json:"-"`                 //time for when added to DB
	Title     string    `json:"title"`             //movie title
	Year      int32     `json:"year,omitempty"`    //movie year
	Runtime   Runtime   `json:"runtime,omitempty"` //movie runtime
	Genres    []string  `json:"genres,omitempty"`  //slices of genres
	Version   int32     `json:"version"`           //increment when updated
}

func ValidateMovie(v *validator.Validator, movie *Movie) {
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")

	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}
