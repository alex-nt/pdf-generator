package collector

import (
	"log"
	"os"
	"path/filepath"

	"github.com/alex-nt/pdf-generator/logger"
)

// PdfStructure is the collection of data needed to generate a pdf
type PdfStructure struct {
	TableOfContents *TOC
	Images          []PdfImage
}

var supportedImageTypes = map[string]bool{
	"jpeg": true,
	"jpg":  true,
	"png":  true,
	"webp": true,
	"gif":  true}

// Gather collects all image data needed for layouting from a folder
func Gather(path string) []PdfStructure {
	fi, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	pdfStructures := make([]PdfStructure, 0)

	switch mode := fi.Mode(); {
	case mode.IsDir():
		readDirectory(path, &pdfStructures)
	case mode.IsRegular():
		log.Fatal("Only directories of images supported!")
	}

	return pdfStructures
}

func readDirectory(path string, pdfStructures *[]PdfStructure) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	pdfImages := make([]PdfImage, 0)

	var tableOfContents *TOC
	for _, f := range files {
		fileName := f.Name()

		newPath := filepath.Join(path, fileName)
		if f.IsDir() {
			readDirectory(newPath, pdfStructures)
		} else {
			ext := filepath.Ext(fileName)[1:]
			name := fileName[0:(len(fileName) - len(ext) - 1)]
			if supportedImageTypes[ext] {
				height, width := size(newPath)

				pdfImages = append(pdfImages, PdfImage{
					Name:   name,
					Height: height,
					Width:  width,
					Path:   newPath,
					Type:   ext})
			} else if fileName == "toc.json" {
				tableOfContents, err = readTOC(newPath)
				if nil != err {
					panic(err)
				}
			} else {
				logger.Info.Printf("File %s not supported", newPath)
			}
		}
	}

	if len(pdfImages) > 0 {
		*pdfStructures = append(*pdfStructures, PdfStructure{Images: pdfImages,
			TableOfContents: tableOfContents})
	}
}
