package decks

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"sync"
)

// Decks is a set of Serato decks that are playing or have played tracks.
// Typically there will be 2 or more decks.
type Decks struct {
	logger *log.Logger

	sync.Mutex
	decks    map[int]*Deck
	snapshot *SessionSnapshot
}

// All returns all the known decks.
func (d *Decks) All() map[int]Deck {
	d.Lock()
	defer d.Unlock()

	decks := make(map[int]Deck, 0)

	for _, deck := range d.decks {
		decks[deck.ID] = *deck
	}

	return decks
}

// Notify will notify each deck with a list of the tracks from the session, so
// the deck can update its own status. Will create any new decks that don't
// exist.
func (d *Decks) Notify(snapshot *SessionSnapshot) error {
	d.Lock()
	defer d.Unlock()

	d.logger.Println("getting new or updated tracked since last snapshot")

	tracks := snapshot.Tracks()

	if d.snapshot != nil {
		tracks = snapshot.NewOrUpdatedTracks(d.snapshot)
	}

	for _, track := range tracks {
		deckID := track.Adat.Deck.Value()

		if _, ok := d.decks[deckID]; !ok {
			d.logger.Printf("creating deck %d\n", deckID)
			d.decks[deckID] = NewDeck(deckID, d.logger)
		}
	}

	for _, deck := range d.decks {
		d.logger.Printf("notifying deck %d\n", deck.ID)

		if err := deck.Notify(tracks); err != nil {
			return err
		}
	}

	d.snapshot = snapshot

	return nil
}

func (d *Decks) String() string {
	d.Lock()
	defer d.Unlock()

	ids := make([]int, 0, len(d.decks))

	for deckID := range d.decks {
		ids = append(ids, deckID)
	}

	sort.Ints(ids)

	var b strings.Builder

	for _, deckID := range ids {
		deck := d.decks[deckID]

		b.WriteString(fmt.Sprintf("deck %d: [%-7v]", deckID, deck.Status))

		if deck.Status == New || deck.Status == Playing {
			b.WriteString(fmt.Sprintf(" %s - %s", deck.Current.Artist(), deck.Current.Title()))
		}

		if deck.Status == Played {
			track := deck.History[len(deck.History)-1]
			b.WriteString(fmt.Sprintf(" %s - %s", track.Artist(), track.Title()))
		}

		b.WriteString("\n")
	}

	return strings.TrimSuffix(b.String(), "\n")
}

// NewDecks returns a new Decks model, with no existing decks.
func NewDecks(logger *log.Logger) *Decks {
	return &Decks{
		logger: logger,
		decks:  make(map[int]*Deck, 0),
	}
}
