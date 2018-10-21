package decks

// Decks is a set of Serato decks that are playing or have played tracks.
// Typically there will be 2 or more decks.
type Decks struct{}

// New returns a new set of Decks that model the currently playing tracks.
func New() *Decks {
	return &Decks{}
}
