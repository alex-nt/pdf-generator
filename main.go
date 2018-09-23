package main

import (
	"flag"
	"os"

	"github.com/alex-nt/pdf-converter/file"
	"github.com/alex-nt/pdf-converter/pdf"
)

func main() {

	directory := flag.String("directory", "", "Directory of images sorted by name. (Required)")
	aspectRatio := flag.Bool("aspectRatio", true, "Keep aspect ratio for files.")

	flag.Parse()

	if "" == *directory {
		flag.PrintDefaults()
		os.Exit(1)
	}

	imageOptions := file.Read(*directory)

	for _, value := range imageOptions {
		pdf.Write(*directory, &value, *aspectRatio)
	}
}
