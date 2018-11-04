package serato

import (
	"bytes"
	"io/ioutil"

	"github.com/tombell/go-serato/serato/chunk"
)

// Session is a current or past session from Serato.
type Session struct {
	Vrsn *chunk.Vrsn
	Oent []*chunk.Oent
	Oren []*chunk.Oren
}

// ReadSession reads and parses the given Serato session file.
func ReadSession(filepath string) (*Session, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	session := &Session{
		Vrsn: nil,
		Oent: make([]*chunk.Oent, 0),
		Oren: make([]*chunk.Oren, 0),
	}

	buf := bytes.NewBuffer(data)
	for buf.Len() > 0 {
		h, err := chunk.NewHeader(buf)
		if err != nil {
			return nil, err
		}

		switch h.Type() {
		case "vrsn":
			vrsn, err := chunk.NewVrsnChunk(h, buf)
			if err != nil {
				return nil, err
			}
			session.Vrsn = vrsn
		case "oent":
			oent, err := chunk.NewOentChunk(h, buf)
			if err != nil {
				return nil, err
			}
			session.Oent = append(session.Oent, oent)
		case "oren":
			oren, err := chunk.NewOrenChunk(h, buf)
			if err != nil {
				return nil, err
			}
			session.Oren = append(session.Oren, oren)
		}
	}

	return session, nil
}
