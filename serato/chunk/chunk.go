package chunk

import (
	"errors"
)

const (
	vrsnID = "vrsn"
	oentID = "oent"
	adatID = "adat"
	orenID = "oren"
	uentID = "uent"
)

// ErrUnexpectedIdentifier is an error representing that a constructor received
// the wrong chunk identifier for the chunk type being created.
var ErrUnexpectedIdentifier = errors.New("unexpected chunk identifier")
