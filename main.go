package main

import (
	"flag"
	"log"
	"os"
	"sync"

	"github.com/alex-nt/pdf-generator/collector"
	"github.com/alex-nt/pdf-generator/logger"
	"github.com/alex-nt/pdf-generator/pdfwriter"
)

func main() {
	directory := flag.String("directory", "", "Directory of images sorted by name.")
	outputDirectory := flag.String("outputDirectory", "", "Output directory for pdfs.")
	jpgOnly := flag.Bool("jpgOnly", true, "Convert all images to jpg.")
	verbose := flag.Bool("v", false, "Verbose mode.")
	marginTopBottom := flag.Int("marginTopBottom", 5, "Top and bottom minimum margin size (can be larger after aspect ratio adjustments)")
	marginLeftRight := flag.Int("marginLeftRight", 5, "Left and right minimum margin size (can be larger after aspect ratio adjustments)")

	flag.Parse()

	logger.Init(*verbose)

	var err error
	workingDirectory := *directory
	if workingDirectory == "" {
		workingDirectory, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	}

	outDirectory := workingDirectory
	if *outputDirectory != "" {
		outDirectory = *outputDirectory
	}

	pdfOptions := pdfwriter.Options{
		Directory:       outDirectory,
		MarginTopBottom: *marginTopBottom,
		MarginLeftRight: *marginLeftRight,
		JPGOnly:         *jpgOnly}
	pdfStructures := collector.Gather(workingDirectory)

	var wg sync.WaitGroup
	for _, pdfStructure := range pdfStructures {
		wg.Add(1)
		go func(pdfStructure collector.PdfStructure, pdfOptions pdfwriter.Options) {
			if err := pdfwriter.Write(pdfStructure, pdfOptions); nil != err {
				panic(err)
			}
			wg.Done()
		}(pdfStructure, pdfOptions)
	}
	wg.Wait()
}
