package main

import (
	"flag"
	"os"

	"github.com/alex-nt/pdf-generator/collector"
	"github.com/alex-nt/pdf-generator/logger"
	"github.com/alex-nt/pdf-generator/pdfwriter"
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

	pdfOptions := pdfwriter.Options{Directory: outDirectory, AspectRatio: *aspectRatio, JPGOnly: *jpgOnly}
	pdfStructures := collector.Gather(*directory)

	for _, pdfStructure := range pdfStructures {
		if err := pdfwriter.Write(pdfStructure, pdfOptions); nil != err {
			panic(err)
		}
	}
}
