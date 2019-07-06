package file

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// PdfImage contains the information needed to add an image to a pdf file
type PdfImage struct {
	Height int
	Width  int
	Path   string
	Name   string
	Type   string
}

// EncodeJPG image to jpg
func (pdfImage *PdfImage) EncodeJPG() bool {
	if pdfImage.Type == "webp" {
		pdfImage.Type = "jpg"
		pdfImage.Path = webpToJPG(pdfImage.Path)
		return true
	}

	if pdfImage.Type == "png" {
		pdfImage.Type = "jpg"
		pdfImage.Path = pngToJPG(pdfImage.Path)
		return true
	}

	return false
}

// TOC is the Table of Contents
// Contains image name indexes
type TOC struct {
	Title    string       `json:"title"`
	Chapters []TOCChapter `json:"chapters"`
}

// TOCChapter is a chapter in the TOC
type TOCChapter struct {
	ChapterName string            `json:"chapterName"`
	Entries     map[string]string `json:"entries"`
}

// PdfStructure is the collection of data needed to generate a pdf
type PdfStructure struct {
	TableOfContents *TOC
	Images          []PdfImage
}

// ReadTOC will read a TOC from a json file
func ReadTOC(path string) (*TOC, error) {
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
