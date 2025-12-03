package data

import "time"

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"` // When movie was added to DB
	Title     string    `json:"title"`
	Year      int       `json:"year,omitempty"`    // Release year
	Runtime   int32     `json:"runtime,omitempty"` // Runtime in minutes
	Genres    []string  `json:"genres,omitempty"`
	Version   int       `json:"version"` // Version number starts at 1 and increments each time movie info is updated
}
