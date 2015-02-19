package bible

import (
	"encoding/json"
	"fmt"
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
		{id: InvalidBook, name: "InvalidBook"},
		{id: 100, name: "BookID(100)"},
	}
	for _, data := range bookNameData {
		if data.id.String() != data.name {
			t.Errorf(
				"Name of BookID %d is %s, expected %s",
				data.id, data.id, data.name)
		}
	}
}

type bookWrapper struct {
	Book BookID
}

func TestJSON(t *testing.T) {
	for b := FirstBook; b <= LastBook; b++ {
		source := bookWrapper{b}
		dest := bookWrapper{}
		data, err := json.Marshal(source)
		if err != nil {
			t.Error(err)
			continue
		}
		err = json.Unmarshal(data, &dest)
		if err != nil {
			t.Error(err)
			continue
		}
		if source.Book != dest.Book {
			t.Errorf(
				"BookID changed after transfering to JSON: %v v.s. %v",
				source, dest)
		}
	}
}

func TestInvalidMarshalJSON(t *testing.T) {
	invalidBookID := []BookID{InvalidBook, -1, 100}
	for _, b := range invalidBookID {
		wrapper := bookWrapper{b}
		data, err := json.Marshal(wrapper)
		if err == nil {
			t.Errorf(
				"Invalid BookID %d is unexpected to be marshaled as %s",
				b, data)
		}
	}
}

func TestInvalidUnmarshalJSON(t *testing.T) {
	invalidBookName := []string{"InvalidBook", "", "BookID(10)"}
	for _, name := range invalidBookName {
		data := []byte(fmt.Sprintf(`{ "Book": "%s" }`, name))
		wrapper := bookWrapper{}
		err := json.Unmarshal(data, &wrapper)
		if err == nil {
			t.Errorf(
				"Invalid book name %s is unexpected to be unmarshaled as %s",
				name, wrapper)
		}
	}
}
