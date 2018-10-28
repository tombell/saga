package decks

import (
	"fmt"
	"sort"
	"strings"
)

// Decks is a set of Serato decks that are playing or have played tracks.
// Typically there will be 2 or more decks.
type Decks struct {
	Snapshot *SessionSnapshot
	Decks    map[int]*Deck
}

// Notify will notify each deck with a list of the tracks from the session, so
// the deck can update its own status. Will create any new decks that don't
// exist.
func (d *Decks) Notify(tracks Tracks) error {
	for _, track := range tracks {
		deckID := track.Adat.Deck.Value()

		if _, ok := d.Decks[deckID]; !ok {
			d.Decks[deckID] = NewDeck(deckID)
		}
	}

	for _, deck := range d.Decks {
		if err := deck.Notify(tracks); err != nil {
			return err
		}
	}

	return nil
}

func (d *Decks) String() string {
	ids := make([]int, 0, len(d.Decks))

	for deckID := range d.Decks {
		ids = append(ids, deckID)
	}

	sort.Ints(ids)

	var b strings.Builder

	for _, deckID := range ids {
		deck := d.Decks[deckID]

		b.WriteString(fmt.Sprintf("Deck %d: [%-7v]", deckID, deck.Status))

		if deck.Status == New || deck.Status == Playing {
			b.WriteString(fmt.Sprintf(" %s - %s", deck.Current.Artist(), deck.Current.Title()))
		}

		if deck.Status == Played {
			track := deck.History[len(deck.History)-1]
			b.WriteString(fmt.Sprintf(" %s - %s", track.Artist(), track.Title()))
		}
	}

	return b.String()
}

// NewDecks returns a new Decks model, with no existing decks.
func NewDecks() *Decks {
	return &Decks{
		Decks: make(map[int]*Deck, 0),
	}
}
