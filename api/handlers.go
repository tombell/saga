package api

import (
	"encoding/json"
	"net/http"

	"github.com/tombell/saga/decks"
)

type deckStatus struct {
	Decks []deckStatusDeck `json:"decks"`
}

type deckStatusDeck struct {
	ID      int               `json:"id"`
	Current *deckStatusTrack  `json:"current"`
	History []deckStatusTrack `json:"history"`
}

type deckStatusTrack struct {
	Status string `json:"status"`
	Artist string `json:"artist"`
	Title  string `json:"title"`
}

func (s *Server) handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := buildStatusResponse(s.decks.All())

		b, err := json.Marshal(status)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))
		w.Header().Add("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		w.Write(b)
	}
}

func buildStatusResponse(decks map[int]decks.Deck) deckStatus {
	status := deckStatus{
		Decks: make([]deckStatusDeck, 0, len(decks)),
	}

	for _, deck := range decks {
		d := deckStatusDeck{
			ID:      deck.ID,
			History: make([]deckStatusTrack, 0, len(deck.History)),
		}

		if deck.Current != nil {
			d.Current = &deckStatusTrack{
				Status: deck.Current.Status().String(),
				Artist: deck.Current.Artist(),
				Title:  deck.Current.Title(),
			}
		}

		for _, track := range deck.History {
			d.History = append(d.History, deckStatusTrack{
				Status: track.Status().String(),
				Artist: track.Artist(),
				Title:  track.Title(),
			})
		}

		status.Decks = append(status.Decks, d)
	}

	return status
}
