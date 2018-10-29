package web

import (
	"fmt"
	"net/http"
)

func (s *Server) handleStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decks := s.decks.All()
		fmt.Fprintf(w, "%#v", decks)
	}
}

func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		fmt.Fprintf(w, template)
	}
}
