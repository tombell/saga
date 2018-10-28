package web

import (
	"fmt"
	"net/http"
)

func (s *Server) handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%v", s.decks)
	}
}
