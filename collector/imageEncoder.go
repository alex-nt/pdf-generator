package collector

import (
	"os"
	"path/filepath"

	"image/gif"
	"image/jpeg"
	"image/png"

	"golang.org/x/image/webp"
)

var quality = jpeg.Options{Quality: 100}

func pngToJPG(path string) string {
	file, err := os.Open(path)
	if nil != err {
		panic(err)
	}
	defer file.Close()

	pngImg, err := png.Decode(file)
	if nil != err {
		panic(err)
	}

	currentExtension := filepath.Ext(path)
	newFilePath := path[0:len(path)-len(currentExtension)] + ".jpg"

	newFile, err := os.Create(newFilePath)
	err = jpeg.Encode(newFile, pngImg, &quality)
	if nil != err {
		panic(err)
	}
	defer newFile.Close()

	return newFilePath
}

func gifToJPG(path string) string {
	file, err := os.Open(path)
	if nil != err {
		panic(err)
	}
	defer file.Close()

	pngImg, err := gif.Decode(file)
	if nil != err {
		panic(err)
	}

	currentExtension := filepath.Ext(path)
	newFilePath := path[0:len(path)-len(currentExtension)] + ".jpg"

	newFile, err := os.Create(newFilePath)
	err = jpeg.Encode(newFile, pngImg, &quality)
	if nil != err {
		panic(err)
	}
	defer newFile.Close()

	return newFilePath
}

func webpToJPG(path string) string {
	file, err := os.Open(path)
	if nil != err {
		panic(err)
	}
	defer file.Close()

	webpImg, err := webp.Decode(file)
	if nil != err {
		panic(err)
	}

	currentExtension := filepath.Ext(path)
	newFilePath := path[0:len(path)-len(currentExtension)] + ".jpg"

	newFile, err := os.Create(newFilePath)
	err = jpeg.Encode(newFile, webpImg, &quality)
	if nil != err {
		panic(err)
	}
	defer newFile.Close()

	return newFilePath
}
