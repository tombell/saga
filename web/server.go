package web

import (
	"net/http"

	"github.com/tombell/saga/decks"
)

// Server is the server that serves the status of the decks for the Serato
// session.
type Server struct {
	mux   *http.ServeMux
	decks *decks.Decks
}

// Run sets up the server handlers and listens on the host/port.
func (s *Server) Run(listen string) error {
	s.mux.HandleFunc("/", s.handler())
	return http.ListenAndServe(listen, s.mux)
}

// NewServer returns a new Server for the given decks from the Serato session.
func NewServer(decks *decks.Decks) *Server {
	return &Server{
		mux:   http.NewServeMux(),
		decks: decks,
	}
}
