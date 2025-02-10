package types

import (
	"time"
)

type Movie struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Genre       string `json:"genre"`
}

type Theater struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type Showtime struct {
	Id             string    `json:"id,omitempty"`
	Date           time.Time `json:"date,omitempty"`
	Movie          Movie     `json:"movie,omitempty"`
	TheaterId      string    `json:"theater_id,omitempty"`
	AvailableSeats string    `json:"available_seats,omitempty"`
	Price          int       `json:"price,omitempty"`
}

type MovieReservation struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Err struct {
	Error string `json:"error"`
}
