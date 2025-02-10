package api

import (
	"net/http"
)

func (s *Server) HandleDeleteMovie(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := s.storage.DeleteMovie(id)
	if err != nil {
		s.respError(w, "Error deleting movie with ID: "+id, 500)
		return
	}
	//TODO: Respond with statusOK and json
}
