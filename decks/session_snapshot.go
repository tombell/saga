package decks

import (
	"github.com/tombell/saga/serato"
)

// SessionSnapshot is a snapshot of the current Serato session.
type SessionSnapshot struct {
	*serato.Session
}
