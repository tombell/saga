package field

import "errors"

// ErrUnexpectedIdentifier is an error representing that a constructor received
// the wrong field identifier for the field type being created.
var ErrUnexpectedIdentifier = errors.New("unexpected field identifier")

// Field ...
type Field interface{}

// Fields ...
type Fields struct {
	Row       *Row       // Field 1
	FullPath  *FullPath  // Field 2
	Location  *Location  // Field 3
	Filename  *Filename  // Field 4
	Title     *Title     // Field 6
	Artist    *Artist    // Field 7
	Album     *Album     // Field 8
	Genre     *Genre     // Field 9
	Length    *Length    // Field 10
	Size      *Size      // Field 11
	Bitrate   *Bitrate   // Field 13
	Frequency *Frequency // Field 14
	BPM       *BPM       // Field 15
	Comment   *Comment   // Field 17
	Language  *Language  // Field 18
	Grouping  *Grouping  // Field 19
	Remixer   *Remixer   // Field 20
	Label     *Label     // Field 21
	Composer  *Composer  // Field 22
	Year      *Year      // Field 23
	StartTime *StartTime // Field 28
	EndTime   *EndTime   // Field 29
	Deck      *Deck      // Field 31
	PlayTime  *PlayTime  // Field 45
	SessionID *SessionID // Field 48
	Played    *Played    // Field 50
	Key       *Key       // Field 51
	Added     *Added     // Field 52
	UpdatedAt *UpdatedAt // Field 53
	Field68   *Field68   // Field 68
	Field69   *Field69   // Field 69
	Field70   *Field70   // Field 70
	Field72   *Field72   // Field 72
}
