package main

import (
	"flag"
	"os"

	"github.com/alex-nt/pdf-converter/file"
	"github.com/alex-nt/pdf-converter/logger"
	"github.com/alex-nt/pdf-converter/pdf"
)

func main() {

	directory := flag.String("directory", "", "Directory of images sorted by name. (Required)")
	outputDirectory := flag.String("outputDirectory", "", "Output directory for pdfs.")
	aspectRatio := flag.Bool("aspectRatio", true, "Preserve image aspect ratio.")
	jpgOnly := flag.Bool("jpgOnly", true, "Convert all images to jpg.")
	verbose := flag.Bool("v", false, "Verbose mode.")

	flag.Parse()

	logger.Init(*verbose)

	if "" == *directory {
		flag.PrintDefaults()
		os.Exit(1)
	}

	outDirectory := *directory
	if "" != *outputDirectory {
		outDirectory = *outputDirectory
	}

	fileOptions := file.Options{JPGOnly: *jpgOnly}
	pdfOptions := pdf.Options{Directory: outDirectory, AspectRatio: *aspectRatio}
	pdfStructures := file.Read(*directory, fileOptions)

	for _, pdfStructure := range pdfStructures {
		pdf.Write(pdfStructure, pdfOptions)
	}
}
