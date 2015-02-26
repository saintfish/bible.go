package bible

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type RefText struct {
	Ref  Ref
	Text string
}

type TextBible interface {
	GetTexts(r RefRange) []RefText
}

var (
	InvalidRefErr = errors.New("Invalid Reference")
	InvalidRef    = Ref{InvalidBook, 0, 0}
)

var textBibleLinePattern = regexp.MustCompile(`^(.+?)(\d+):(\d+)\s+(.*)$`)

func LoadTextBible(filePath string) (TextBible, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	bible := newInmemoryBible()
	book := InvalidBook
	bookText := ""
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		groups := textBibleLinePattern.FindStringSubmatch(scanner.Text())
		if len(groups) != 5 {
			return nil, fmt.Errorf("Malformed line in file %s: %s", filePath, scanner.Text())
		}
		if bookText != groups[1] {
			book++
			bookText = groups[1]
		}
		chapter, err := strconv.Atoi(groups[2])
		if err != nil {
			return nil, fmt.Errorf("Malformed chapter number in file %s: %s", filePath, scanner.Text())
		}
		verse, err := strconv.Atoi(groups[3])
		if err != nil {
			return nil, fmt.Errorf("Malformed chapter number in file %s: %s", filePath, scanner.Text())
		}
		bible.addVerse(Ref{book, chapter, verse}, groups[4])
	}
	bible.doneAddVerse()
	return bible, nil
}

type sortRef []Ref

func (r sortRef) Len() int {
	return len(r)
}

func lessThanRef(a, b Ref) bool {
	if a.Book != b.Book {
		return a.Book < b.Book
	} else if a.Chapter != b.Chapter {
		return a.Chapter < b.Chapter
	} else {
		return a.Verse < b.Verse
	}
}
func (r sortRef) Less(i, j int) bool {
	return lessThanRef(r[i], r[j])
}

func (r sortRef) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type inmemoryBible struct {
	m map[Ref]string
	k sortRef
}

func newInmemoryBible() *inmemoryBible {
	return &inmemoryBible{
		m: make(map[Ref]string),
		k: make(sortRef, 0),
	}
}

func (b *inmemoryBible) GetTexts(rr RefRange) []RefText {
	rt := []RefText{}
	for ref := b.resolveRef(rr.Begin); ref != InvalidRef && (lessThanRef(ref, rr.End) || ref == rr.End); ref = b.nextRef(ref) {
		rt = append(rt, RefText{ref, b.m[ref]})
	}
	return rt
}

func (b *inmemoryBible) resolveRef(r Ref) Ref {
	index := sort.Search(b.k.Len(), func(i int) bool {
		return !lessThanRef(b.k[i], r)
	})
	if index >= b.k.Len() {
		return InvalidRef
	}
	return b.k[index]
}

func (b *inmemoryBible) nextRef(r Ref) Ref {
	if r == InvalidRef {
		return InvalidRef
	}
	index := sort.Search(b.k.Len(), func(i int) bool {
		return !lessThanRef(b.k[i], r) && b.k[i] != r
	})
	if index >= b.k.Len() {
		return InvalidRef
	}
	return b.k[index]
}

func (b *inmemoryBible) addVerse(r Ref, text string) {
	b.k = append(b.k, r)
	b.m[r] = text
}

func (b *inmemoryBible) doneAddVerse() {
	sort.Sort(b.k)
}
