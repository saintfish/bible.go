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
	BookBegin    = 0
	BookEnd      = math.MaxInt16
	ChapterBegin = 0
	ChapterEnd   = math.MaxInt16
)

type RefRange struct {
	Begin, End Ref
}

type RefRangeList []RefRange

func SingleVerseRef(b BookID, c, v int) RefRangeList {
	return RefRangeList{
		RefRange{
			Begin: Ref{b, c, v},
			End:   Ref{b, c, v},
		},
	}
}

func SingleRangeRef(b BookID, c, v1, v2 int) RefRangeList {
	return RefRangeList{
		RefRange{
			Begin: Ref{b, c, v1},
			End:   Ref{b, c, v2},
		},
	}
}
