package api

import (
	"net/http"
)

func (s *Server) HandleGetTheater(w http.ResponseWriter, r *http.Request) {
	theater, err := s.storage.GetTheater(r.PathValue("id"))
	if err != nil {
		s.respError(w, "Not found", http.StatusNotFound)
		return
	}

	s.respondJSON(w, theater, http.StatusOK)
}
