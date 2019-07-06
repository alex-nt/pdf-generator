package file

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Read extracts all image data needed for layouting from a folder
func Read(path string) []PdfStructure {
	fi, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	pdfStructures := make([]PdfStructure, 0, 0)

	switch mode := fi.Mode(); {
	case mode.IsDir():
		readDirectory(path, &pdfStructures)
	case mode.IsRegular():
		log.Fatal("Only directories of images supported!")
	}

	return pdfStructures
}

func readDirectory(path string, pdfStructures *[]PdfStructure) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	pdfImages := make([]PdfImage, 0, 0)

	for _, f := range files {
		newPath := filepath.Join(path, f.Name())
		if f.IsDir() {
			readDirectory(newPath, pdfStructures)
		} else {
			height, width := size(newPath)

			pdfImages = append(pdfImages, PdfImage{
				Height: height,
				Width:  width,
				Path:   newPath,
				Type:   filepath.Ext(f.Name())[1:]})
		}
	}

	if len(pdfImages) > 0 {
		*pdfStructures = append(*pdfStructures, PdfStructure{Images: pdfImages})
	}
}
