package decks

// Decks is a set of Serato decks that are playing or have played tracks.
// Typically there will be 2 or more decks.
type Decks struct {
	decks    map[int]*Deck
	snapshot *SessionSnapshot
}

// Notify will notify each deck with a list of the tracks from the session, so
// the deck can update its own status.
func (d *Decks) Notify(tracks Tracks) {
}

// NewDecks returns a new Decks model, will initialise any decks using the
// initial list of tracks from the session.
func NewDecks(tracks Tracks) *Decks {
	return &Decks{
		decks:    make(map[int]*Deck),
		snapshot: nil,
	}
}
