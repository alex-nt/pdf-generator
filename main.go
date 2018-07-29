package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/alex-nt/pdf-converter/file"
	"github.com/alex-nt/pdf-converter/pdf"
)

func main() {

	outputName := flag.String("outputName", "", "Output file name.")
	directory := flag.String("directory", "", "Directory of images sorted by name. (Required)")
	marginTop := flag.Float64("marginTop", 0, "Top and bottom margins.")
	marginLeft := flag.Float64("marginLeft", 0, "Left and right margins.")
	aspectRatio := flag.Bool("aspectRatio", false, "Keep aspect ratio for files.")

	flag.Parse()

	if "" == *directory {
		flag.PrintDefaults()
		os.Exit(1)
	}

	imageOptions := file.Read(*directory)

	fmt.Println(*aspectRatio)
	pdf.Write(*outputName, *directory, imageOptions, *marginTop, *marginLeft, *aspectRatio)
}
