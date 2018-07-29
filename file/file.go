package file

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Read extracts all image data needed for layouting from a folder
func Read(path string) *[]ImageInfo {
	fi, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		return readDirectory(path)
	case mode.IsRegular():
		panic("Only directories of images supported!")
	}
	return nil
}

func readDirectory(path string) *[]ImageInfo {
	imageInformation := make([]ImageInfo, 0, 0)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		imagePath := filepath.Join(path, f.Name())
		height, width := size(imagePath)

		extension := filepath.Ext(f.Name())
		imageInformation = append(imageInformation, ImageInfo{
			Height: height,
			Width:  width,
			Path:   imagePath,
			Type:   extension[1:]})
	}

	return &imageInformation
}
