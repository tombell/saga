package decks

import (
	"github.com/tombell/saga/serato"
)

// SessionSnapshot is a snapshot of the current Serato session.
type SessionSnapshot struct {
	*serato.Session
}

// Tracks returns a map of tracks where the map key is the row for the track.
func (s *SessionSnapshot) Tracks() map[int]Track {
	tracks := make(map[int]Track, 0)

	for _, oent := range s.Oent {
		track := NewTrack(*oent.Adat)
		tracks[track.Row.Value()] = *track
	}

	for _, oren := range s.Oren {
		delete(tracks, oren.Uent.Value())
	}

	return tracks
}
