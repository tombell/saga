package field

import "errors"

// ErrUnexpectedIdentifier is an error representing that a constructor received
// the wrong field identifier for the field type being instantiated.
var ErrUnexpectedIdentifier = errors.New("invalid field identifier")

// Field ...
type Field interface{}

// Fields ...
type Fields struct {
	Row       *Row       // Field 1
	FullPath  *FullPath  // Field 4
	Title     *Title     // Field 6
	Artist    *Artist    // Field 7
	Genre     *Genre     // Field 9
	BPM       *BPM       // Field 15
	Grouping  *Grouping  // Field 19
	Label     *Label     // Field 21
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

	// TODO:
	// Album
	// Bitrate
	// Comment
	// Composer
	// Filename
	// Frequency
	// Language
	// Length
	// Location
	// Remixer
	// Size
}
