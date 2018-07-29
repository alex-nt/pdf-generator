package main

import (
	"flag"
	"os"

	"github.com/alex-nt/pdf-converter/file"
	"github.com/alex-nt/pdf-converter/pdf"
)

func main() {

	outputName := flag.String("outputName", "", "Output file name. (Required)")
	directory := flag.String("directory", "", "Directory of images sorted by name. (Required)")
	marginTop := flag.Float64("marginTop", 0, "Top and bottom margins.")
	marginLeft := flag.Float64("marginLeft", 0, "Left and right margins.")
	aspectRatio := flag.Bool("aspectRatio", false, "Keep aspect ratio for files.")

	flag.Parse()

	if "" == *outputName || "" == *directory {
		flag.PrintDefaults()
		os.Exit(1)
	}

	imageOptions := file.Read(*directory)

	pdf.Write(*outputName, imageOptions, *marginTop, *marginLeft, *aspectRatio)
}
