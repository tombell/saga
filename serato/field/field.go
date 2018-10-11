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

// Fields is a struct containing all possible fields for a single ADAT chunk.
type Fields struct {
	Row       *Row
	FullPath  *FullPath
	Location  *Location
	Filename  *Filename
	Title     *Title
	Artist    *Artist
	Album     *Album
	Genre     *Genre
	Length    *Length
	Size      *Size
	Bitrate   *Bitrate
	Frequency *Frequency
	BPM       *BPM
	Comment   *Comment
	Language  *Language
	Grouping  *Grouping
	Remixer   *Remixer
	Label     *Label
	Composer  *Composer
	Year      *Year
	StartTime *StartTime
	EndTime   *EndTime
	Deck      *Deck
	PlayTime  *PlayTime
	SessionID *SessionID
	Played    *Played
	Key       *Key
	Added     *Added
	UpdatedAt *UpdatedAt
	Field68   *Field68
	Field69   *Field69
	Field70   *Field70
	Field72   *Field72
}
