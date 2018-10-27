package decks

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
func (d *Deck) Notify(tracks Tracks) {
}

// NewDeck returns a new Deck with an initial empty state.
func NewDeck(id int) *Deck {
	return &Deck{
		ID:      id,
		Status:  Empty,
		Current: nil,
		History: make([]*Track, 0),
		maxRow:  0,
	}
}
