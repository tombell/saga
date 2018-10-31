package server

import (
	"log"
	"net/http"

	"github.com/tombell/saga/decks"
)

// Config contains configuration and data needed to run the server.
type Config struct {
	Logger  *log.Logger
	Decks   *decks.Decks
	Address string
}

// Server serves the status of the decks for the Serato session.
type Server struct {
	logger  *log.Logger
	decks   *decks.Decks
	address string
	mux     *http.ServeMux
}

// Run sets up the server handlers and listens on the host/port.
func (s *Server) Run(ch chan error) {
	s.mux.HandleFunc("/", s.handler())

	if err := http.ListenAndServe(s.address, s.mux); err != nil {
		ch <- err
	}
}

// New returns a new Server for the given decks from the Serato session.
func New(cfg Config) *Server {
	return &Server{
		logger:  cfg.Logger,
		decks:   cfg.Decks,
		address: cfg.Address,
		mux:     http.NewServeMux(),
	}
}
