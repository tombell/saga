package serato

import "github.com/tombell/saga/serato/chunk"

// Session is a current or past session from Serato.
type Session struct {
	Vrsn *chunk.Vrsn
	Oent []*chunk.Oent
	Oren []*chunk.Oren
}

// ReadSession reads and parses the given Serato session file.
func ReadSession(filepath string) (*Session, error) {
	return nil, nil
}
