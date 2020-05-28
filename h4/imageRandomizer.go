// https://www.socketloop.com/tutorials/golang-get-rgba-values-of-each-image-pixel
// https://stackoverflow.com/questions/33186783/get-a-pixel-array-from-from-golang-image-image
// https://socketloop.com/tutorials/golang-convert-integer-to-binary-octal-hexadecimal-and-back-to-integer
// https://dev.to/andyhaskell/how-i-made-a-slick-personal-logo-with-go-s-standard-library-29j9
package main

import (
	"fmt"
	"flag"
	"image"
	"image/color"
	"image/png"
	// ^ https://golang.org/pkg/image/
	"os"
)

func decodeImg(myImage string) image.Image {
	file, err := os.Open(myImage)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return data
}

func recolor(data image.Image, bounds image.Rectangle, number int) *image.RGBA {
	recolored := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
		colors := data.At(x, y)
		r, g, b, a := colors.RGBA()
		// Convert RGBA values to 8-bit
		r = ((r / 257) * uint32(number)) % 255
		g = ((g / 257) * uint32(number)) % 255
		b = ((b / 257) * uint32(number)) % 255
		newColor := color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
		recolored.Set(x, y, newColor)
		}
	}
	return recolored
}

func main() {

	var image string
	var output string
	var number int

	flag.StringVar(&image, "f", "", "PNG File to use")
	flag.StringVar(&output, "o", "", "Output file")
	flag.IntVar(&number, "n", 3, "Randomizer number")
	flag.Parse()

	// Decode image into something useable and define bounds
	data := decodeImg(image)
	bounds := data.Bounds()
	
	// Create an outfile
	out, err := os.Create(output)
	if err != nil {
		fmt.Println(err)
	}
	defer out.Close()
	
	// Make the new image
	recolored := recolor(data, bounds, number)
	if err := png.Encode(out, recolored); err != nil {
		fmt.Println(err)
	}
}