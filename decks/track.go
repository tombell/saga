package decks

import (
	"github.com/tombell/saga/serato/chunk"
)

// Track represents a new, playing, played, or skipped track on a deck in
// Serato.
type Track struct {
	chunk.Adat
}

// Status returns the tracks current status, whether it's playing, played,
// new, or skipped.
func (t *Track) Status() Status {
	if t.Fields.Played.Value() {
		if t.Fields.PlayTime != nil {
			return Played
		}

		return Playing
	}

	if t.Fields.PlayTime != nil {
		return Skipped
	}

	return New
}

// Artist returns the artist of the track.
func (t *Track) Artist() string {
	if t.Adat.Artist == nil {
		return ""
	}

	return t.Adat.Artist.Value()
}

// Title returns the title of the track.
func (t *Track) Title() string {
	if t.Adat.Title == nil {
		return ""
	}

	return t.Adat.Title.Value()
}

// NewTrack returns an initialised Track using the given ADAT chunk.
func NewTrack(adat chunk.Adat) *Track {
	return &Track{
		Adat: adat,
	}
}
