package main

import (
	"github.com/alex-nt/pdf-converter/config"
	"github.com/alex-nt/pdf-converter/file"
	"github.com/alex-nt/pdf-converter/pdf"
)

func main() {
	imageOptions := file.Read("C:\\Users\\alexa\\Downloads\\Over The Garden Wall 016 (2017) (Digital-Empire)")

	profile := config.Read("C:\\Migration\\Projects\\pdf-converter\\profile.json")

	pdf.Write("test.pdf", imageOptions, &profile)
}
