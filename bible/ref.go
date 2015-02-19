package bible

import (
	"math"
)

type Ref struct {
	Book    BookID
	Chapter int
	Verse   int
}

const (
	BookBegin  = 0
	BookEnd    = math.MaxInt16
	VerseBegin = 0
	VerseEnd   = math.MaxInt16
)

type RefRange struct {
	Begin, End Ref
}

type RefRangeList []RefRange
