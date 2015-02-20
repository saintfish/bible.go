package bible

import (
	"testing"
)

func TestBookNameListFromFile(t *testing.T) {
	files := []string{
		"../data/zh-hans.bible.unv/book_names.json",
		"../data/zh-hant.bible.unv/book_names.json",
	}
	for _, fn := range files {
		_, err := BookNameListFromFile(fn)
		if err != nil {
			t.Error(err)
		}
	}
}
