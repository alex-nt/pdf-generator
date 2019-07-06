package pdf

import (
	"os"
	"strings"

	gofpdf "github.com/jung-kurt/gofpdf/v2"

	"github.com/alex-nt/pdf-converter/file"
	"github.com/alex-nt/pdf-converter/logger"
)

func Write(directory string, imageDetails *[]file.ImageInfo, aspectRatio bool) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	width, height := pdf.GetPageSize()
	sizeType := gofpdf.SizeType{Wd: width, Ht: height}

	for _, image := range *imageDetails {
		orientation := pageOrientation(image)

		pdf.AddPageFormat(orientation, sizeType)

		addImage(pdf, image, aspectRatio)
	}

	outputFileName := generateName((*imageDetails)[0].Path, directory)
	err := pdf.OutputFileAndClose(outputFileName)
	if nil != err {
		panic(err)
	}
}

func generateName(name string, directory string) string {
	parts := strings.Split(name, string(os.PathSeparator))

	return directory + string(os.PathSeparator) + parts[len(parts)-2] + ".pdf"
}

func addImage(pdf *gofpdf.Fpdf, imageDetails file.ImageInfo, aspectRatio bool) {
	var opt gofpdf.ImageOptions
	opt.ImageType = imageDetails.Type

	width, height := pdf.GetPageSize()

	var marginLeft, marginTop float64
	if aspectRatio {
		computedHeight := (float64(imageDetails.Height) / float64(imageDetails.Width)) * width
		if height < computedHeight {
			computedWidth := (float64(imageDetails.Width) / float64(imageDetails.Height)) * height
			marginLeft = (width - computedWidth) / 2
			width = computedWidth
		} else {
			height = computedHeight
			marginTop = (height - computedHeight) / 2
		}
	}

	logger.Debug.Printf("Type: %s, Path: %s, Img w: %d h: %d, Output Img w: %f h: %f \n", imageDetails.Type, imageDetails.Path, imageDetails.Width, imageDetails.Height, width, height)
	pdf.ImageOptions(imageDetails.Path, marginLeft, marginTop,
		width, height, false, opt, 0, "")
}

func pageOrientation(image file.ImageInfo) string {
	orientation := "P"
	if image.Height < image.Width {
		orientation = "L"
	}

	return orientation
}
