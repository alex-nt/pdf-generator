package pdf

import (
	"github.com/jung-kurt/gofpdf"

	"github.com/alex-nt/pdf-converter/file"
)

func Write(name string, imageDetails *[]file.ImageInfo, marginTop, marginLeft float64, aspectRatio bool) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	width, height := pdf.GetPageSize()
	sizeType := gofpdf.SizeType{
		Wd: width, Ht: height}

	for _, image := range *imageDetails {
		orientation := pageFormat(image)

		pdf.AddPageFormat(orientation, sizeType)

		addImage(pdf, image, marginTop, marginLeft)
	}

	err := pdf.OutputFileAndClose(name)
	if nil != err {
		panic(err)
	}
}

func addImage(pdf *gofpdf.Fpdf, imageDeails file.ImageInfo, marginTop, marginLeft float64) {
	var opt gofpdf.ImageOptions
	opt.ImageType = imageDeails.Type

	width, height := pdf.GetPageSize()

	pdf.ImageOptions(imageDeails.Path, marginTop, marginLeft,
		width-2*marginTop, height-2*marginLeft, false, opt, 0, "")
}

func pageFormat(image file.ImageInfo) string {
	orientation := "P"
	if image.Height < image.Width {
		orientation = "L"
	}

	return orientation
}
