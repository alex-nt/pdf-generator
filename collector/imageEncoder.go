package collector

import (
	"bufio"
	"bytes"
	"os"

	"image/gif"
	"image/jpeg"
	"image/png"

	"golang.org/x/image/webp"
)

var quality = jpeg.Options{Quality: 100}

func pngToJPG(path string) bytes.Buffer {
	file, err := os.Open(path)
	if nil != err {
		panic(err)
	}
	defer file.Close()

	pngImg, err := png.Decode(file)
	if nil != err {
		panic(err)
	}

	var buffer bytes.Buffer
	err = jpeg.Encode(bufio.NewWriter(&buffer), pngImg, &quality)
	if nil != err {
		panic(err)
	}

	return buffer
}

func gifToJPG(path string) bytes.Buffer {
	file, err := os.Open(path)
	if nil != err {
		panic(err)
	}
	defer file.Close()

	gifImg, err := gif.Decode(file)
	if nil != err {
		panic(err)
	}

	var buffer bytes.Buffer
	err = jpeg.Encode(bufio.NewWriter(&buffer), gifImg, &quality)
	if nil != err {
		panic(err)
	}

	return buffer
}

func webpToJPG(path string) bytes.Buffer {
	file, err := os.Open(path)
	if nil != err {
		panic(err)
	}
	defer file.Close()

	webpImg, err := webp.Decode(file)
	if nil != err {
		panic(err)
	}

	var buffer bytes.Buffer
	err = jpeg.Encode(bufio.NewWriter(&buffer), webpImg, &quality)
	if nil != err {
		panic(err)
	}

	return buffer
}
