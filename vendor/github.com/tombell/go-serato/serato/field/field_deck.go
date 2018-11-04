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
func (f *Deck) Value() int {
	return int(binary.BigEndian.Uint32(f.data))
}

func (f *Deck) String() string {
	return fmt.Sprintf("%d", f.Value())
}

// NewDeckField returns an initialised Deck, using the given field header.
func NewDeckField(header *Header, r io.Reader) (*Deck, error) {
	if header.Identifier != deckID {
		return nil, ErrUnexpectedIdentifier
	}

	data := make([]byte, header.Length)
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return nil, err
	}

	return &Deck{
		header: header,
		data:   data[:],
	}, nil
}
