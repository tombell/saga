package decks

// Status is the status that a deck could be in.
type Status int

func (s Status) String() string {
	switch s {
	case Empty:
		return "EMPTY"
	case New:
		return "NEW"
	case Playing:
		return "PLAYING"
	case Played:
		return "PLAYED"
	case Skipped:
		return "SKIPPED"
	}

	return "INVALID"
}

// All valid deck statuses.
const (
	Empty Status = iota
	New
	Playing
	Played
	Skipped
)
