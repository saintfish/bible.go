package bible

import (
	"encoding/json"
	"fmt"
)

//go:generate stringer --type BookID

// BookID is a numerical identifier of a book in the bible.
type BookID int

// A list of books in the bible
const (
	// InvalidBook is an invalid identifier.
	InvalidBook BookID = iota

	Genesis
	Exodus
	Leviticus
	Numbers
	Deuteronomy
	Joshua
	Judges
	Ruth
	Samuel1
	Samuel2
	Kings1
	Kings2
	Chronicles1
	Chronicles2
	Ezra
	Nehemiah
	Esther
	Job
	Psalm
	Proverbs
	Ecclesiastes
	SongOfSongs
	Isaiah
	Jeremiah
	Lamentations
	Ezekiel
	Daniel
	Hosea
	Joel
	Amos
	Obadiah
	Jonah
	Micah
	Nahum
	Habakkuk
	Zephaniah
	Haggai
	Zechariah
	Malachi
	Matthew
	Mark
	Luke
	John
	Acts
	Romans
	Corinthians1
	Corinthians2
	Galatians
	Ephesians
	Philippians
	Colossians
	Thessalonians1
	Thessalonians2
	Timothy1
	Timothy2
	Titus
	Philemon
	Hebrews
	James
	Peter1
	Peter2
	John1
	John2
	John3
	Jude
	Revelation
)

// A few alias useful in looping through the books
const (
	FirstBook BookID = Genesis
	LastBook         = Revelation
	FirstOT          = Genesis
	LastOT           = Malachi
	FirstNT          = Matthew
	LastNT           = Revelation
)

// NumBooks is the number of the books in the bible
const NumBooks = int(LastBook - FirstBook + 1)

var nameBookIDMap map[string]BookID

func init() {
	nameBookIDMap = make(map[string]BookID)
	for b := FirstBook; b <= LastBook; b++ {
		nameBookIDMap[b.String()] = b
	}
}

func BookIDFromName(name string) BookID {
	b, ok := nameBookIDMap[name]
	if !ok {
		return InvalidBook
	}
	return b
}

func (b BookID) MarshalJSON() ([]byte, error) {
	if b >= FirstBook && b <= LastBook {
		return json.Marshal(b.String())
	}
	return nil, fmt.Errorf("Invalid BookID %s", b)
}

func (b *BookID) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	book := BookIDFromName(s)
	if book == InvalidBook {
		return fmt.Errorf("Invalid book name %s", s)
	}
	*b = book
	return nil
}
