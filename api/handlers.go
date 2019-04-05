package api

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func (s *Server) handler() http.HandlerFunc {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			// TODO: error logging
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		go s.register(conn)
	}
}
