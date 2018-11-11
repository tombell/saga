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
	bitrateID   = 13
	bpmID       = 15
	commentID   = 17
	groupingID  = 19
	remixerID   = 20
	labelID     = 21
	composerID  = 22
	yearID      = 23
	starttimeID = 28
	endtimeID   = 29
	deckID      = 31
	field39ID   = 39
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

// ErrUnexpectedIdentifier is an error returned when a field constructor is
// given the wrong field identifier for the field type being created.
var ErrUnexpectedIdentifier = errors.New("unexpected field identifier")
