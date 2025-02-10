package api

import (
	"encoding/json"
	"net/http"

	"github.com/hacdan/cinema-api/types"
)

func (s *Server) HandlePutMovie(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	movie := types.Movie{}

	err := decoder.Decode(&movie)
	if err != nil {
		s.respError(w, "Error decoding", 500) //TODO: Change to HTTP Status Error from http package
		return
	}

	updatedMovie, err := s.storage.UpdateMovie(movie)
	if err != nil {
		s.respError(w, "Error updating movie", 500)
		return
	}
	s.respondJSON(w, updatedMovie, http.StatusOK)
}
