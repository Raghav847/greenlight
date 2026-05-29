package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`                //unique int ID for the movie
	CreatedAt time.Time `json:"-"`                 //time for when added to DB
	Title     string    `json:"title"`             //movie title
	Year      int32     `json:"year,omitempty"`    //movie year
	Runtime   Runtime   `json:"runtime,omitempty"` //movie runtime
	Genres    []string  `json:"genres,omitempty"`  //slices of genres
	Version   int32     `json:"version"`           //increment when updated
}
