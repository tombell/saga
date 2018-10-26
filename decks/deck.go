package decks

// Deck is a single deck in Serato that can play a track.
type Deck struct {
	ID     int
	Status Status

	Current  *Track
	Previous *Track
}

// Update ...
func (d *Deck) Update(tracks map[int]Track) error {
	return nil
}

// NewDeck returns a new Deck, that has a current and previous track.
func NewDeck(id int) *Deck {
	return &Deck{
		ID:     id,
		Status: Empty,
	}
}
