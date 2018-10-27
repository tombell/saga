package decks

import (
	"errors"
	"sort"
)

var (
	// ErrUnknownTransitionFromStatus is an error when a track will transition
	// from an unknown status.
	ErrUnknownTransitionFromStatus = errors.New("unknown transition from status")

	// ErrUnknownTransitionToStatus is an error when a track will transition to
	// an unknown status.
	ErrUnknownTransitionToStatus = errors.New("unknown transition to status")

	// ErrInvalidTransitionToEmpty is an error when a track will transition to
	// empty.
	ErrInvalidTransitionToEmpty = errors.New("invalid transition to empty")
)

// Deck is a single deck in Serato. The Deck has a currently playing track, and
// history of played tracks.
type Deck struct {
	ID     int
	Status Status

	Current *Track
	History []*Track

	maxRow int
}

// Notify will notify the deck with a list of the tracks from the session. The
// deck will transition the track if the track status has changed.
func (d *Deck) Notify(tracks Tracks) error {
	mine := make(Tracks, 0)
	rows := make([]int, 0)

	for row, track := range tracks {
		if track.Deck.Value() == d.ID && row >= d.maxRow {
			mine[row] = track
			rows = append(rows, row)
		}
	}

	sort.Ints(rows)

	for _, row := range rows {
		track, ok := mine[row]
		if !ok {
			continue
		}

		if err := d.transitionTo(track); err != nil {
			return err
		}
	}

	return nil
}

func (d *Deck) transitionTo(track Track) error {
	from := Empty

	if d.Current != nil {
		from = d.Current.Status()
	}

	switch from {
	case Empty:
		return d.transitionFromEmpty(track)
	case New:
		return d.transitionFromNew(track)
	case Playing:
		return d.transitionFromPlaying(track)
	case Played:
		return d.transitionFromPlayed(track)
	case Skipped:
		return d.transitionFromSkipped(track)
	default:
		return ErrUnknownTransitionFromStatus
	}
}

func (d *Deck) transitionFromEmpty(track Track) error {
	to := track.Status()

	switch to {
	case Empty:
		return ErrInvalidTransitionToEmpty
	case New:
		d.transitionFromEmptyToNew(track)
	case Playing:
		d.transitionFromEmptyToNew(track)
		d.transitionFromNewToPlaying(track)
	case Played:
		d.transitionFromEmptyToNew(track)
		d.transitionFromNewToPlaying(track)
		d.transitionFromPlayingToPlayed(track)
	case Skipped:
		d.transitionFromEmptyToNew(track)
		d.transitionFromNewToSkipped(track)
	default:
		return ErrUnknownTransitionToStatus
	}

	return nil
}

func (d *Deck) transitionFromNew(track Track) error {
	to := track.Status()

	switch to {
	case Empty:
		return ErrInvalidTransitionToEmpty
	case New:
		d.transitionFromNewToSkipped(track)
		d.transitionFromSkippedToNew(track)
	case Playing:
		d.transitionFromNewToPlaying(track)
	case Played:
		d.transitionFromNewToPlaying(track)
		d.transitionFromPlayingToPlayed(track)
	case Skipped:
		d.transitionFromNewToSkipped(track)
	default:
		return ErrUnknownTransitionToStatus
	}

	return nil
}

func (d *Deck) transitionFromPlaying(track Track) error {
	to := track.Status()

	switch to {
	case Empty:
		return ErrInvalidTransitionToEmpty
	case New:
		d.transitionFromPlayingToPlayed(track)
		d.transitionFromPlayedToNew(track)
	case Playing:
		d.transitionFromPlayingToPlayed(track)
		d.transitionFromPlayedToNew(track)
		d.transitionFromNewToPlaying(track)
	case Played:
		d.transitionFromPlayingToPlayed(track)
	case Skipped:
		d.transitionFromNewToSkipped(track)
	default:
		return ErrUnknownTransitionToStatus
	}

	return nil
}

func (d *Deck) transitionFromPlayed(track Track) error {
	to := track.Status()

	switch to {
	case Empty:
		return ErrInvalidTransitionToEmpty
	case New:
		d.transitionFromPlayedToNew(track)
	case Playing:
		d.transitionFromPlayedToNew(track)
		d.transitionFromNewToPlaying(track)
	case Played:
		d.transitionFromPlayedToNew(track)
		d.transitionFromNewToPlaying(track)
		d.transitionFromPlayingToPlayed(track)
	case Skipped:
		d.transitionFromPlayedToNew(track)
		d.transitionFromNewToSkipped(track)
	default:
		return ErrUnknownTransitionToStatus
	}

	return nil
}

func (d *Deck) transitionFromSkipped(track Track) error {
	to := track.Status()

	switch to {
	case Empty:
		return ErrInvalidTransitionToEmpty
	case New:
		d.transitionFromSkippedToNew(track)
	case Playing:
		d.transitionFromSkippedToNew(track)
		d.transitionFromNewToPlaying(track)
	case Played:
		d.transitionFromSkippedToNew(track)
		d.transitionFromNewToSkipped(track)
	case Skipped:
		d.transitionFromSkippedToNew(track)
		d.transitionFromNewToPlaying(track)
		d.transitionFromPlayingToPlayed(track)
	default:
		return ErrUnknownTransitionToStatus
	}

	return nil
}

func (d *Deck) transitionFromEmptyToNew(track Track) {
	d.maxRow = max(track.Row.Value(), d.maxRow)
	d.Current = &track
	d.Status = track.Status()
}

func (d *Deck) transitionFromPlayedToNew(track Track) {
	d.transitionFromEmptyToNew(track)
}

func (d *Deck) transitionFromSkippedToNew(track Track) {
	d.transitionFromEmptyToNew(track)
}

func (d *Deck) transitionFromNewToPlaying(track Track) {
	d.maxRow = max(track.Row.Value(), d.maxRow)
	d.Current = &track
	d.Status = track.Status()
}

func (d *Deck) transitionFromNewToSkipped(track Track) {
	d.maxRow = max(track.Row.Value(), d.maxRow)
	d.Current = nil
	d.Status = track.Status()
}

func (d *Deck) transitionFromPlayingToPlayed(track Track) {
	d.maxRow = max(track.Row.Value(), d.maxRow)
	d.Current = nil
	d.Status = track.Status()
	d.History = append(d.History, &track)
}

// NewDeck returns a new Deck with an initial empty state.
func NewDeck(id int) *Deck {
	return &Deck{
		ID:      id,
		Status:  Empty,
		History: make([]*Track, 0),
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
