package file

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
