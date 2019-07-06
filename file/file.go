package file

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Options is a collection of settings to be applied on the files
type Options struct {
	JPGOnly bool
}

// Read extracts all image data needed for layouting from a folder
func Read(path string, fileOptions Options) []PdfStructure {
	fi, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}

	pdfStructures := make([]PdfStructure, 0, 0)

	switch mode := fi.Mode(); {
	case mode.IsDir():
		readDirectory(path, &pdfStructures, fileOptions)
	case mode.IsRegular():
		log.Fatal("Only directories of images supported!")
	}

	return pdfStructures
}

func readDirectory(path string, pdfStructures *[]PdfStructure, fileOptions Options) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	pdfImages := make([]PdfImage, 0, 0)

	for _, f := range files {

		newPath := filepath.Join(path, f.Name())
		if f.IsDir() {
			readDirectory(newPath, pdfStructures, fileOptions)
		} else {
			height, width := size(newPath)

			extension, imagePath, transcodedImage := processImage(newPath, fileOptions)

			pdfImages = append(pdfImages, PdfImage{
				Height:          height,
				Width:           width,
				Path:            imagePath,
				Type:            extension,
				DeleteAfterUser: transcodedImage})
		}
	}

	if len(pdfImages) > 0 {
		*pdfStructures = append(*pdfStructures, PdfStructure{Images: pdfImages})
	}
}

func processImage(path string, options Options) (string, string, bool) {
	extension := filepath.Ext(path)[1:]

	transcodedImage := false
	if options.JPGOnly {
		if extension == "webp" {
			extension = "jpg"
			transcodedImage = true
			path = webpToJPG(path)
		}

		if extension == "png" {
			extension = "jpg"
			transcodedImage = true
			path = pngToJPG(path)
		}
	}

	return extension, path, transcodedImage
}
