// This example demonstrates decoding a JPEG image and examining its pixels.
package main

import (
	//"encoding/base64"
	"fmt"
	"image"
	"log"
	//"strings"
	"os"

	// Package image/jpeg is not used explicitly in the code below,
	// but is imported for its initialization side-effect, which allows
	// image.Decode to understand JPEG formatted images. Uncomment these
	// two lines to also understand GIF and PNG images:
	// _ "image/gif"
	// "image/png"
	_ "image/jpeg"
)

func main() {
	// Decode the JPEG data. If reading from file, create a reader with
	//
	reader, err := os.Open("img.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	// reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()

	// An image's bounds do not necessarily start at (0, 0), so the two loops start
	// at bounds.Min.Y and bounds.Min.X. Looping over Y first and X second is more
	// likely to result in better memory access patterns than X first and Y second.

	greens := 0

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {

			r, g, b, _ := m.At(x, y).RGBA()
			fmt.Print("x: ", x, " y: ", y, " r: ", r, " g: ", g, " b: ", b)
			//count green pixels
			if g > 10000 && r < 100000 && b < 10000 {
				greens++
				fmt.Println("    GREEN!!!!!")
			} else {
				fmt.Println()
			}

		}
	}
	fmt.Println("greens", greens)
}