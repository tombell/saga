package web

import (
	"net/http"

	"github.com/tombell/saga/decks"
)

// Config ...
type Config struct{}

// Server is the server that serves the status of the decks for the Serato
// session.
type Server struct {
	mux   *http.ServeMux
	decks *decks.Decks
}

// Run sets up the server handlers and listens on the host/port.
func (s *Server) Run(listen string, errCh chan error) {
	s.mux.HandleFunc("/_status", s.handleStatus())
	s.mux.HandleFunc("/", s.handleIndex())

	if err := http.ListenAndServe(listen, s.mux); err != nil {
		errCh <- err
	}
}

// NewServer returns a new Server for the given decks from the Serato session.
func NewServer(decks *decks.Decks) *Server {
	return &Server{
		mux:   http.NewServeMux(),
		decks: decks,
	}
}
