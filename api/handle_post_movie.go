package api

import (
	"encoding/json"
	"net/http"

	"github.com/hacdan/cinema-api/types"
)

func (s *Server) HandlePostMovie(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	movie := types.Movie{}

	err := decoder.Decode(&movie)
	if err != nil {
		s.respError(w, "Error decoding JSON", 500)
		return
	}

	newMovie, err := s.storage.AddMovie(movie)
	if err != nil {
		s.respError(w, "Error creating movie", 500)
		return
	}
	s.respondJSON(w, newMovie, http.StatusOK)
}
