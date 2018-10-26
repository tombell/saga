package decks

import (
	"github.com/tombell/saga/serato/chunk"
)

// Track represents a playing or played track on a deck in Serato.
type Track struct {
	*chunk.Adat
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

// NewTrack ...
func NewTrack(adat *chunk.Adat) *Track {
	return &Track{
		Adat: adat,
	}
}
