package bible

import (
	"testing"
)

func TestBookName(t *testing.T) {
	var bookNameData = []struct {
		id   BookID
		name string
	}{
		{id: FirstBook, name: "Genesis"},
		{id: LastBook, name: "Revelation"},
		{id: FirstOT, name: "Genesis"},
		{id: LastOT, name: "Malachi"},
		{id: FirstNT, name: "Matthew"},
		{id: LastNT, name: "Revelation"},
		{id: 1, name: "Genesis"},
		{id: InvalidBook, name: ""},
		{id: 100, name: ""},
	}
	for _, data := range bookNameData {
		if BookName(data.id) != data.name {
			t.Errorf(
				"Name of BookID %d is %s, expected %s",
				data.id, BookName(data.id), data.name)
		}
	}
}
