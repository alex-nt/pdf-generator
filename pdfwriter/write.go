package pdfwriter

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jung-kurt/gofpdf/v2"

	"github.com/alex-nt/pdf-generator/collector"
	"github.com/alex-nt/pdf-generator/logger"
)

// Options contains the settings for pdf generation
type Options struct {
	Directory       string
	MarginTopBottom int
	MarginLeftRight int
}

// Write will generate and write on disk a pdf
func Write(pdfStructure collector.PdfStructure, options Options) error {
	pdf := gofpdf.New("P", "mm", "A4", "")

	sectionMap := addTOC(pdf, pdfStructure.TableOfContents)

	width, height := pdf.GetPageSize()
	sizeType := gofpdf.SizeType{Wd: width, Ht: height}

	for _, image := range pdfStructure.Images {
		orientation := pageOrientation(image)

		pdf.AddPageFormat(orientation, sizeType)
		addImage(pdf, image, sectionMap, options)
	}

	outputFilePath := generateOutputFilePath(pdfStructure.Images[0].Path, options)
	logger.Info.Println("\tOutput pdf path\t", outputFilePath)
	return pdf.OutputFileAndClose(outputFilePath)
}

type section struct {
	nr     int
	name   string
	linkID int
}

func addTOC(pdf *gofpdf.Fpdf, toc *collector.TOC) map[string]*section {
	sectionMap := make(map[string]*section)

	if nil != toc {
		pdf.AddPage()

		pdf.SetFont("Arial", "B", 16)
		pdf.Cell(40, 20, toc.Title)
		pdf.Ln(20)

		pdf.SetFont("Arial", "", 12)
		for idx, entry := range toc.Entries {
			generatedLinkID := pdf.AddLink()
			sectionMap[entry.File] = &section{
				name:   entry.Name,
				nr:     idx,
				linkID: generatedLinkID}

			pdf.WriteLinkID(10, generateTOCLine(pdf, entry.Name, idx), generatedLinkID)
			pdf.Ln(10)
		}
	}

	return sectionMap
}

func generateTOCLine(pdf *gofpdf.Fpdf, name string, idx int) string {
	contentLength := pdf.GetStringWidth(name) + pdf.GetStringWidth(strconv.Itoa(idx))

	width, _ := pdf.GetPageSize()
	left, _, right, _ := pdf.GetMargins()

	lineWidth := width - contentLength - left - right
	separatorLen := pdf.GetStringWidth(".")

	nrOfDots := int(lineWidth/separatorLen) - 10

	return fmt.Sprintf("%s%s{%d}", name, strings.Repeat(".", nrOfDots), idx)
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

func addImage(pdf *gofpdf.Fpdf, image collector.PdfImage, sections map[string]*section, options Options) {
	section := sections[image.Name]
	if nil != section {
		pdf.RegisterAlias(fmt.Sprintf("{%d}", section.nr), fmt.Sprintf("%d", pdf.PageNo()))
		pdf.SetLink(section.linkID, -1, pdf.PageNo())
	}
	var opt gofpdf.ImageOptions
	opt.ImageType = image.Type

	width, height := pdf.GetPageSize()

	imageLayout := computeImageSize(image, width, height, options)
 
	logger.Debug.Printf("Image %v", image)
	// All images will be written as JPG due to weird DPT-RP1 behaviour
	pdf.RegisterImageReader(image.Name, "jpg", image.Reader())
	pdf.ImageOptions(image.Name, imageLayout.MarginLeft, imageLayout.MarginTop,
		imageLayout.Width, imageLayout.Height, false, opt, 0, "")
}

func computeImageSize(image collector.PdfImage, pageWidth float64, pageHeight float64, options Options) imageLayout {
	var marginLeft, marginTop float64

	effectivePageHeight := pageHeight - float64(2*options.MarginTopBottom)
	effectivePageWidth := pageWidth - float64(2*options.MarginLeftRight)

	computedHeight := (float64(image.Height) / float64(image.Width)) * effectivePageWidth
	if effectivePageHeight < computedHeight {
		computedWidth := (float64(image.Width) / float64(image.Height)) * effectivePageHeight
		marginLeft = (effectivePageWidth - computedWidth) / 2
		effectivePageWidth = computedWidth
	} else {
		effectivePageHeight = computedHeight
		marginTop = (effectivePageHeight - computedHeight) / 2
	}

	return imageLayout{Height: effectivePageHeight,
		Width:      effectivePageWidth,
		MarginTop:  marginTop + float64(options.MarginTopBottom),
		MarginLeft: marginLeft + float64(options.MarginLeftRight)}
}

func pageOrientation(image collector.PdfImage) string {
	orientation := "P"
	if image.Height < image.Width {
		orientation = "L"
	}

	return orientation
}
