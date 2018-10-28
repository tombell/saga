package web

import (
	"fmt"
	"net/http"
)

func (s *Server) handleStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.decks.Lock()
		decks := s.decks.Decks
		s.decks.Unlock()

		fmt.Fprintf(w, "%#v", decks)
	}
}

func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		fmt.Fprintf(w, template)
	}
}
