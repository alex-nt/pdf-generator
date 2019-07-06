package collector

import (
	"image"
	"os"

	"github.com/alex-nt/pdf-generator/logger"
)

// PdfImage contains the information needed to add an image to a pdf file
type PdfImage struct {
	Height int
	Width  int
	Path   string
	Name   string
	Type   string
}

// EncodeJPG encodes image to jpg
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

func size(path string) (height, width int) {
	logger.Debug.Println(path)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		logger.Debug.Println(path)
		panic(err)
	}
	return image.Height, image.Width
}
