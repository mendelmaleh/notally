package notally

import (
	"encoding/xml"
	"strconv"
	"time"
)

// somewhat generated with https://github.com/miku/zek

type Date struct {
	time.Time
}

func (d *Date) UnmarshalText(data []byte) error {
	msec, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}

	d.Time = time.UnixMilli(msec)

	return nil
}

// Common fields of notes and lists
type Common struct {
	// Text        string `xml:",chardata"`
	Color       string `xml:"color"`
	DateCreated Date   `xml:"date-created"`
	Pinned      string `xml:"pinned"`
	Title       string `xml:"title"`
}

// Note is a regular note
type Note struct {
	Common
	Body string `xml:"body"`
}

// List is a list note
type List struct {
	Common
	Item []struct {
		// Chardata string `xml:",chardata"`
		Text    string `xml:"text"`
		Checked string `xml:"checked"`
	} `xml:"item"`
}

// Notes contains regular notes and lists
type Notes struct {
	// Text string `xml:",chardata"`
	Note []Note `xml:"note"`
	List []List `xml:"list"`
}

// ExportedNotes represents a Notally export
type ExportedNotes struct {
	XMLName xml.Name `xml:"exported-notes"`
	// Text    string   `xml:",chardata"`

	Notes         Notes `xml:"notes"`
	ArchivedNotes Notes `xml:"archived-notes"`
	DeletedNotes  Notes `xml:"deleted-notes"`

	Label []string `xml:"label"`
}
