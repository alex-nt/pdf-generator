package collector

import (
	"bufio"
	"image"
	"io"
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

func (pdfImage *PdfImage) Reader() io.Reader {
	switch pdfImage.Type {
	case "jpg", "jpeg":
		file, err := os.Open(pdfImage.Path)
		if err != nil {
			panic(err)
		}
		return file
	case "png":
		buffer := pngToJPG(pdfImage.Path)
		return bufio.NewReader(&buffer)
	case "webp":
		buffer := webpToJPG(pdfImage.Path)
		return bufio.NewReader(&buffer)
	default:
		panic("Type " + pdfImage.Type + " not supported!")
	}
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
