package collector

import (
	"os"
	"path/filepath"

	"image/jpeg" // jpeg decoder
	"image/png"

	"golang.org/x/image/webp" // webp decoder
)

func pngToJPG(path string) string {
	file, err := os.Open(path)
	if nil != err {
		panic(err)
	}

	pngImg, err := png.Decode(file)
	if nil != err {
		panic(err)
	}

	currentExtension := filepath.Ext(path)
	newFilePath := path[0:len(path)-len(currentExtension)] + ".jpg"

	newFile, err := os.Create(newFilePath)
	err = jpeg.Encode(newFile, pngImg, nil)
	if nil != err {
		panic(err)
	}

	return newFilePath
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
