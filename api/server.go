package api

import (
	"net/http"

	"github.com/hacdan/cinema-api/storage"
)

type Server struct {
	listenAddr string
	storage    storage.Storage
}

func NewServer(listenAddr string, dbURL string) Server {
	storage := storage.New(dbURL)

	return Server{
		listenAddr: listenAddr,
		storage:    storage,
	}
}

func (s *Server) Start() error {
	// Movie Endpoints
	http.HandleFunc("GET /v1/movies", s.HandleGetMovies)
	http.HandleFunc("GET /v1/movie", s.HandleGetMovie)
	http.HandleFunc("POST /v1/movie", s.HandlePostMovie)
	http.HandleFunc("PUT /v1/movie", s.HandlePutMovie)
	http.HandleFunc("DELETE /v1/movie/{id}", s.HandleDeleteMovie)

	// Theater Endpoints
	http.HandleFunc("GET /v1/theater", s.HandleGetTheater)
	http.HandleFunc("GET /v1/theaters", s.HandleGetTheaters)
	http.HandleFunc("GET /v1/teather/showtimes", s.HandleGetShowTimes)

	http.HandleFunc("POST /v1/reservation", s.HandlePostReservation)
	http.HandleFunc("POST /v1/reservation/confirmation", s.HandlePostReservationConfirmation)
	http.HandleFunc("GET /v1/reservation/{id}", s.HandleGetReservation)

	return http.ListenAndServe(s.listenAddr, nil)
}
