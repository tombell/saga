package field

import "errors"

const (
	rowID       = 1
	fullpathID  = 2
	locationID  = 3
	filenameID  = 4
	titleID     = 6
	artistID    = 7
	albumID     = 8
	genreID     = 9
	lengthID    = 10
	sizeID      = 11
	bitrateID   = 13
	frequencyID = 14
	bpmID       = 15
	commentID   = 17
	languageID  = 18
	groupingID  = 19
	remixerID   = 20
	labelID     = 21
	composerID  = 22
	yearID      = 23
	starttimeID = 28
	endtimeID   = 29
	deckID      = 31
	playtimeID  = 45
	sessionID   = 48
	playedID    = 50
	keyID       = 51
	addedID     = 52
	updatedAtID = 53
	field68ID   = 68
	field69ID   = 69
	field70ID   = 70
	field72ID   = 72
)

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
