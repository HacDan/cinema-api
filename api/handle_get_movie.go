package api

import (
	"net/http"
)

func (s *Server) HandleGetMovie(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	movie, err := s.storage.GetMovie(id)
	if err != nil {
		panic(err)
	}
	s.respondJSON(w, movie, http.StatusOK)
}
