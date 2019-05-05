package file

import (
	"fmt"
	"image"
	"os"
	"path/filepath"

	// jpeg decoder
	"image/jpeg"
	_ "image/png" // png decoder

	"golang.org/x/image/webp" // webp decoder
)

func size(path string) (height, width int) {
	fmt.Println(path)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		panic(err)
	}
	return image.Height, image.Width
}

func webpToJPG(path string) string {
	file, err := os.Open(path)
	if nil != err {
		panic(err)
	}

	webpImg, err := webp.Decode(file)
	if nil != err {
		panic(err)
	}

	currentExtension := filepath.Ext(path)
	newFilePath := path[0:len(path)-len(currentExtension)] + ".jpg"

	newFile, err := os.Create(newFilePath)
	err = jpeg.Encode(newFile, webpImg, nil)
	if nil != err {
		panic(err)
	}

	return newFilePath
}
