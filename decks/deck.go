package decks

// Status is the status that a deck could be in.
type Status int

// All valid deck statuses.
const (
	Empty Status = iota
	New
	Playing
	Played
	Skipped
)

// Deck is a single deck in Serato that can play a track.
type Deck struct {
	ID     int
	Status Status
}

// NewDeck returns a new Deck.
func NewDeck(id int) *Deck {
	return &Deck{id, Empty}
}
