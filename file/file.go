package file

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Read extracts all image data needed for layouting from a folder
func Read(path string) [][]ImageInfo {
	fi, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	imagePerDir := make([][]ImageInfo, 0, 0)

	switch mode := fi.Mode(); {
	case mode.IsDir():
		readDirectory(path, &imagePerDir)
	case mode.IsRegular():
		panic("Only directories of images supported!")
	}

	return imagePerDir
}

func readDirectory(path string, imagePerDir *[][]ImageInfo) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	imageInformation := make([]ImageInfo, 0, 0)

	for _, f := range files {

		newPath := filepath.Join(path, f.Name())
		if f.IsDir() {
			readDirectory(newPath, imagePerDir)
		} else {
			height, width := size(newPath)

			extension := filepath.Ext(f.Name())[1:]
			if extension == "webp" {
				newPath = webpToJPG(newPath)
				extension = filepath.Ext(newPath)[1:]
			}

			imageInformation = append(imageInformation, ImageInfo{
				Height: height,
				Width:  width,
				Path:   newPath,
				Type:   extension})
		}
	}

	if len(imageInformation) > 0 {
		*imagePerDir = append(*imagePerDir, imageInformation)
	}
}
