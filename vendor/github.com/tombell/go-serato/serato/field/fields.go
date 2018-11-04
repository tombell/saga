package field

import (
	"bytes"
)

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
	Field39   *Field39
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

// NewFields returns an initialised Fields struct after reading the field data
// from the given bytes buffer.
func NewFields(data []byte) (*Fields, error) {
	fields := &Fields{}

	buf := bytes.NewBuffer(data)
	for buf.Len() > 0 {
		h, err := NewHeader(buf)
		if err != nil {
			return nil, err
		}

		switch h.Identifier {
		case 1:
			field, err := NewRowField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Row = field
		case 2:
			field, err := NewFullPathField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.FullPath = field
		case 3:
			field, err := NewLocationField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Location = field
		case 4:
			field, err := NewFilenameField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Filename = field
		case 6:
			field, err := NewTitleField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Title = field
		case 7:
			field, err := NewArtistField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Artist = field
		case 8:
			field, err := NewAlbumField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Album = field
		case 9:
			field, err := NewGenreField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Genre = field
		case 10:
			field, err := NewLengthField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Length = field
		case 11:
			field, err := NewSizeField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Size = field
		case 13:
			field, err := NewBitrateField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Bitrate = field
		case 14:
			field, err := NewFrequencyField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Frequency = field
		case 15:
			field, err := NewBPMField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.BPM = field
		case 17:
			field, err := NewCommentField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Comment = field
		case 18:
			field, err := NewLanguageField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Language = field
		case 19:
			field, err := NewGroupingField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Grouping = field
		case 20:
			field, err := NewRemixerField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Remixer = field
		case 21:
			field, err := NewLabelField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Label = field
		case 22:
			field, err := NewComposerField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Composer = field
		case 23:
			field, err := NewYearField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Year = field
		case 28:
			field, err := NewStartTimeField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.StartTime = field
		case 29:
			field, err := NewEndTimeField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.EndTime = field
		case 31:
			field, err := NewDeckField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Deck = field
		case 39:
			field, err := NewField39Field(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Field39 = field
		case 45:
			field, err := NewPlayTimeField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.PlayTime = field
		case 48:
			field, err := NewSessionIDField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.SessionID = field
		case 50:
			field, err := NewPlayedField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Played = field
		case 51:
			field, err := NewKeyField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Key = field
		case 52:
			field, err := NewAddedField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Added = field
		case 53:
			field, err := NewUpdatedAtField(h, buf)
			if err != nil {
				return nil, err
			}
			fields.UpdatedAt = field
		case 68:
			field, err := NewField68Field(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Field68 = field
		case 69:
			field, err := NewField69Field(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Field69 = field
		case 70:
			field, err := NewField70Field(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Field70 = field
		case 72:
			field, err := NewField72Field(h, buf)
			if err != nil {
				return nil, err
			}
			fields.Field72 = field
		default:
			buf.Next(int(h.Length))
		}
	}

	return fields, nil
}
