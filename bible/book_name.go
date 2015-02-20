package bible

import (
	"encoding/json"
	"io/ioutil"
)

type BookName struct {
	ID      BookID
	Full    string
	Abbr    string
	Aliases []string
}

type BookNameList []BookName

func BookNameListFromFile(fn string) (BookNameList, error) {
	content, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	bookNames := BookNameList{}
	err = json.Unmarshal(content, &bookNames)
	if err != nil {
		return nil, err
	}
	return bookNames, nil
}
