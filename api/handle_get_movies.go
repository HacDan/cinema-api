package api

import (
	"net/http"
)

func (s *Server) HandleGetMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := s.storage.GetMovies()
	if err != nil {
		panic(err) //TODO: handle error response
	}
	s.respondJSON(w, movies, http.StatusOK)
}
