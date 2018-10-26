package decks

// Deck is a single deck in Serato that can play a track.
type Deck struct {
	ID     int
	Status Status

	Current  *Track
	Previous *Track
}

// Notify will notify the deck with a list of the tracks from the session. The
// deck will transition the track if the track status has changed.
func (d *Deck) Notify(tracks Tracks) {
}

// NewDeck returns a new Deck, that has a current and previous track.
func NewDeck(id int) *Deck {
	return &Deck{
		ID:     id,
		Status: Empty,
	}
}
