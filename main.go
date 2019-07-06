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
	aspectRatio := flag.Bool("aspectRatio", true, "Preserve image aspect ratio.")
	//jpgOnly := flag.Bool("jpg-only", true, "Convert all images to jpg.")
	verbose := flag.Bool("v", false, "Verbose mode.")

	flag.Parse()

	logger.Init(*verbose)

	if "" == *directory {
		flag.PrintDefaults()
		os.Exit(1)
	}

	imageOptions := file.Read(*directory)

	for _, value := range imageOptions {
		pdf.Write(*directory, &value, *aspectRatio)
	}
}
