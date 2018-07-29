package file

import (
	"fmt"
	"image"
	"os"

	_ "image/jpeg" // jpeg decoder
	_ "image/png"  // png decoder
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
