package collector

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// TOC is the Table of Contents
// Contains image name indexes
type TOC struct {
	Title   string       `json:"title"`
	Entries []TOCChapter `json:"entries"`
}

// TOCChapter is a chapter in the TOC
type TOCChapter struct {
	Name string `json:"name"`
	File string `json:"file"`
}

func readTOC(path string) (*TOC, error) {
	jsonFile, err := os.Open(path)
	if nil != err {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if nil != err {
		return nil, err
	}

	var tableOfContents TOC

	if err = json.Unmarshal(byteValue, &tableOfContents); nil != err {
		return nil, err
	}
	return &tableOfContents, nil
}
