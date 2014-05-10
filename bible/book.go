package bible

// BookID is a numerical identifier of a book in the bible.
type BookID byte

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
	FirstBook = Genesis
	LastBook  = Revelation
	FirstOT   = Genesis
	LastOT    = Malachi
	FirstNT   = Matthew
	LastNT    = Revelation
)

// NumBooks is the number of the books in the bible
const NumBooks = int(LastBook - FirstBook + 1)

var bookNames = [...]string{
	"", // invalid
	"Genesis",
	"Exodus",
	"Leviticus",
	"Numbers",
	"Deuteronomy",
	"Joshua",
	"Judges",
	"Ruth",
	"Samuel1",
	"Samuel2",
	"Kings1",
	"Kings2",
	"Chronicles1",
	"Chronicles2",
	"Ezra",
	"Nehemiah",
	"Esther",
	"Job",
	"Psalm",
	"Proverbs",
	"Ecclesiastes",
	"SongOfSongs",
	"Isaiah",
	"Jeremiah",
	"Lamentations",
	"Ezekiel",
	"Daniel",
	"Hosea",
	"Joel",
	"Amos",
	"Obadiah",
	"Jonah",
	"Micah",
	"Nahum",
	"Habakkuk",
	"Zephaniah",
	"Haggai",
	"Zechariah",
	"Malachi",
	"Matthew",
	"Mark",
	"Luke",
	"John",
	"Acts",
	"Romans",
	"Corinthians1",
	"Corinthians2",
	"Galatians",
	"Ephesians",
	"Philippians",
	"Colossians",
	"Thessalonians1",
	"Thessalonians2",
	"Timothy1",
	"Timothy2",
	"Titus",
	"Philemon",
	"Hebrews",
	"James",
	"Peter1",
	"Peter2",
	"John1",
	"John2",
	"John3",
	"Jude",
	"Revelation",
}

// BookName gives the string representation of BookID enum.
func BookName(b BookID) string {
	if b >= FirstBook && b <= LastBook {
		return bookNames[b]
	}
	return bookNames[InvalidBook]
}
