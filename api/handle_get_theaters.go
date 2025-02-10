package api

import (
	"net/http"
)

func (s *Server) HandleGetTheaters(w http.ResponseWriter, r *http.Request) {
	theaters, err := s.storage.GetTheaters()
	if err != nil {
		s.respError(w, "Error: Theaters not found", 500)
		return
	}

	s.respondJSON(w, theaters, http.StatusOK)
}
