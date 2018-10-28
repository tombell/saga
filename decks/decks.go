package decks

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

// Decks is a set of Serato decks that are playing or have played tracks.
// Typically there will be 2 or more decks.
type Decks struct {
	mu       sync.Mutex
	Snapshot *SessionSnapshot
	decks    map[int]*Deck
}

// Notify will notify each deck with a list of the tracks from the session, so
// the deck can update its own status. Will create any new decks that don't
// exist.
func (d *Decks) Notify(tracks Tracks) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	for _, track := range tracks {
		deckID := track.Adat.Deck.Value()

		if _, ok := d.decks[deckID]; !ok {
			d.decks[deckID] = NewDeck(deckID)
		}
	}

	for _, deck := range d.decks {
		if err := deck.Notify(tracks); err != nil {
			return err
		}
	}

	return nil
}

// TODO: this will need cleaning up...
func (d *Decks) String() string {
	var b strings.Builder

	ids := make([]int, 0)

	for deckID := range d.decks {
		ids = append(ids, deckID)
	}

	sort.Ints(ids)

	for _, deckID := range ids {
		deck := d.decks[deckID]

		b.WriteString(fmt.Sprintf("Deck %d: [%-7v]", deckID, deck.Status))

		if deck.Status == New || deck.Status == Playing {
			b.WriteString(fmt.Sprintf(" %s - %s", deck.Current.Artist(), deck.Current.Title()))
		}

		if deck.Status == Played {
			track := deck.History[len(deck.History)-1]
			b.WriteString(fmt.Sprintf(" %s - %s", track.Artist(), track.Title()))
		}

		b.WriteString("\n")
	}

	return b.String()
}

// NewDecks returns a new Decks model, with no existing decks.
func NewDecks() *Decks {
	return &Decks{
		decks: make(map[int]*Deck, 0),
	}
}
