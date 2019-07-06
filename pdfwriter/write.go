package pdfwriter

import (
	"os"
	"strings"

	gofpdf "github.com/jung-kurt/gofpdf/v2"

	"github.com/alex-nt/pdf-generator/collector"
	"github.com/alex-nt/pdf-generator/logger"
)

// Options contains the settings for pdf generation
type Options struct {
	Directory   string
	AspectRatio bool
	JPGOnly     bool
}

// Write will generate and write on disk a pdf
func Write(pdfStructure collector.PdfStructure, options Options) error {
	pdf := gofpdf.New("P", "mm", "A4", "")

	width, height := pdf.GetPageSize()
	sizeType := gofpdf.SizeType{Wd: width, Ht: height}

	for _, image := range pdfStructure.Images {
		orientation := pageOrientation(image)
		if options.JPGOnly {
			if image.EncodeJPG() {
				defer deleteImage(image)
			}
		}
		pdf.AddPageFormat(orientation, sizeType)
		addImage(pdf, image, options)
	}

	outputFilePath := generateOutputFilePath(pdfStructure.Images[0].Path, options)
	logger.Info.Println("\tOutput pdf path\t", outputFilePath)
	return pdf.OutputFileAndClose(outputFilePath)
}

func deleteImage(image collector.PdfImage) {
	if err := os.Remove(image.Path); nil != err {
		panic(err)
	}
}

func generateOutputFilePath(name string, options Options) string {
	parts := strings.Split(name, string(os.PathSeparator))
	logger.Info.Println(options.Directory)

	if err := os.MkdirAll(options.Directory, os.ModePerm); nil != err {
		panic(err)
	}

	pdfPath := options.Directory + string(os.PathSeparator) + parts[len(parts)-2] + ".pdf"
	return pdfPath
}

type imageLayout struct {
	Height     float64
	Width      float64
	MarginTop  float64
	MarginLeft float64
}

func addImage(pdf *gofpdf.Fpdf, image collector.PdfImage, options Options) {
	var opt gofpdf.ImageOptions
	opt.ImageType = image.Type

	width, height := pdf.GetPageSize()

	imageLayout := computeImageSize(image, width, height, options)

	logger.Debug.Printf("Image %v", image)
	pdf.ImageOptions(image.Path, imageLayout.MarginLeft, imageLayout.MarginTop,
		imageLayout.Width, imageLayout.Height, false, opt, 0, "")
}

func computeImageSize(image collector.PdfImage, width float64, height float64, options Options) imageLayout {
	var marginLeft, marginTop float64

	if options.AspectRatio {
		computedHeight := (float64(image.Height) / float64(image.Width)) * width
		if height < computedHeight {
			computedWidth := (float64(image.Width) / float64(image.Height)) * height
			marginLeft = (width - computedWidth) / 2
			width = computedWidth
		} else {
			height = computedHeight
			marginTop = (height - computedHeight) / 2
		}
	}

	return imageLayout{Height: height,
		Width:      width,
		MarginTop:  marginTop,
		MarginLeft: marginLeft}
}

func pageOrientation(image collector.PdfImage) string {
	orientation := "P"
	if image.Height < image.Width {
		orientation = "L"
	}

	return orientation
}
