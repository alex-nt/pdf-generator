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
	aspectRatio := flag.Bool("aspectRatio", true, "Preserve image aspect ratio.")
	jpgOnly := flag.Bool("jpgOnly", true, "Convert all images to jpg.")
	verbose := flag.Bool("v", false, "Verbose mode.")

	flag.Parse()

	logger.Init(*verbose)

	var err error
	workingDirectory := *directory
	if "" == workingDirectory {
		workingDirectory, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	}

	outDirectory := workingDirectory
	if "" != *outputDirectory {
		outDirectory = *outputDirectory
	}

	pdfOptions := pdfwriter.Options{Directory: outDirectory, AspectRatio: *aspectRatio, JPGOnly: *jpgOnly}
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
