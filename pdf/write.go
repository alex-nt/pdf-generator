package pdf

import (
	"github.com/jung-kurt/gofpdf"

	"github.com/alex-nt/pdf-converter/config"
	"github.com/alex-nt/pdf-converter/file"
)

func Write(name string, imageDetails *[]file.ImageInfo, profile *config.OutputProfile) {
	pdf := gofpdf.New("P", "mm", "A4", "")

	width, height := pdf.GetPageSize()
	sizeType := gofpdf.SizeType{
		Wd: width, Ht: height}

	for _, image := range *imageDetails {
		orientation := pageFormat(image, profile)

		pdf.AddPageFormat(orientation, sizeType)

		addImage(pdf, image, profile)
	}

	err := pdf.OutputFileAndClose(name)
	if nil != err {
		panic(err)
	}
}

func addImage(pdf *gofpdf.Fpdf, imageDeails file.ImageInfo, profile *config.OutputProfile) {
	var opt gofpdf.ImageOptions
	opt.ImageType = imageDeails.Type

	width, height := pdf.GetPageSize()

	pdf.ImageOptions(imageDeails.Path, float64(profile.Margins.Top), float64(profile.Margins.Left),
		float64(width), float64(height), false, opt, 0, "")
}

func pageFormat(image file.ImageInfo, profile *config.OutputProfile) string {
	orientation := "P"
	if image.Height < image.Width {
		orientation = "L"
	}

	return orientation
}
