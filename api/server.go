package api

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/tombell/saga/decks"
)

// Config contains configuration and data needed to run the server.
type Config struct {
	Logger *log.Logger
	Decks  *decks.Decks
	Listen string
}

// Server serves the status of the decks for the Serato session.
type Server struct {
	logger *log.Logger
	decks  *decks.Decks
	listen string
	mux    *http.ServeMux
}

// Run sets up the server handlers and listens on the host/port.
func (s *Server) Run(ch chan error) {
	s.mux.HandleFunc("/", s.handler())

	if err := http.ListenAndServe(s.listen, s.mux); err != nil {
		ch <- err
	}
}

func (s *Server) register(conn *websocket.Conn) {
	ch := make(chan bool, 1)

	s.decks.AddNotificationChannel(ch)

	for {
		select {
		case <-ch:
			status := buildStatusResponse(s.decks.All())

			if err := conn.WriteJSON(status); err != nil {
				// TODO: error logging
				break
			}
		}
	}

	s.decks.RemoveNotificationChannel(ch)
	conn.Close()
}

// New returns a new Server for the given decks from the Serato session.
func New(cfg Config) *Server {
	return &Server{
		logger: cfg.Logger,
		decks:  cfg.Decks,
		listen: cfg.Listen,
		mux:    http.NewServeMux(),
	}
}
