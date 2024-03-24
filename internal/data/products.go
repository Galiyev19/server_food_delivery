package data

import "time"

type Rating struct {
	Rate  float64
	Count int64
}

type Products struct {
	ID       int64     `json:"id"`
	CreadAt  time.Time `json:"-"`
	Title    string    `json:"title"`
	Category string    `json:"category"`
	Price    int64     `json:"price"`
	Image    string    `json:"image"`
	Rating   Rating    `json:"rating"`
	Version  int64     `json:"version"`
}
