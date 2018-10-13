package field

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Deck is the deck that the track is playing on in Serato.
type Deck struct {
	header *Header
	data   []byte
}

// Value returns the deck.
func (d *Deck) Value() int {
	return int(binary.BigEndian.Uint32(d.data))
}

func (d *Deck) String() string {
	return fmt.Sprintf("Deck: %d", d.Value())
}

// NewDeckField returns a Deck, using the header to read the field data.
func NewDeckField(header *Header, r io.Reader) (*Deck, error) {
	if header.Identifier != deckID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)

	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Deck{header, data[:]}, nil
}
