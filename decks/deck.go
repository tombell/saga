package decks

import (
	"fmt"
)

// Deck is a single deck in Serato that can play a track.
type Deck struct {
	ID     int
	Status Status

	Current  *Track
	Previous *Track
}

// Update ...
func (d *Deck) Update(newer *SessionSnapshot) error {
	return nil
}

func (d *Deck) String() string {
	return fmt.Sprintf("Deck %d\n", d.ID)
}

// NewDeck returns a new Deck, that has a current and previous track.
func NewDeck(id int) *Deck {
	return &Deck{
		ID:     id,
		Status: Empty,
	}
}
