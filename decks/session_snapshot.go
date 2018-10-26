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

// NewOrUpdatedTracks returns the tracks that are new or have been updated in
// the newer session snapshot when compared to the previous snapshot.
func (s *SessionSnapshot) NewOrUpdatedTracks(older *SessionSnapshot) map[int]Track {
	newerTracks := s.Tracks()
	olderTracks := older.Tracks()

	newerRows := make([]int, len(newerTracks))
	olderRows := make([]int, len(olderTracks))

	for row := range newerTracks {
		newerRows = append(newerRows, row)
	}

	for row := range olderTracks {
		olderRows = append(olderRows, row)
	}

	tracks := make(map[int]Track, 0)
	added := diff(newerRows, olderRows)

	for _, row := range added {
		tracks[row] = newerTracks[row]
	}

	for row, track := range newerTracks {
		older, ok := olderTracks[row]
		if !ok || track.UpdatedAt.Value().After(older.UpdatedAt.Value()) {
			tracks[row] = track
		}
	}

	return tracks
}

func diff(a, b []int) []int {
	var diff []int
	m := make(map[int]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}

	return diff
}
