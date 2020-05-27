// https://www.socketloop.com/tutorials/golang-get-rgba-values-of-each-image-pixel
// https://stackoverflow.com/questions/33186783/get-a-pixel-array-from-from-golang-image-image
// https://socketloop.com/tutorials/golang-convert-integer-to-binary-octal-hexadecimal-and-back-to-integer
// https://dev.to/andyhaskell/how-i-made-a-slick-personal-logo-with-go-s-standard-library-29j9
package main

import (
	"fmt"
	"image"
	_"image/png"
	// ^ https://golang.org/pkg/image/
	"os"
	//"strconv"
	//"io/ioutil"
)

func main() {

	file, err := os.Open("stego.png")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	data, _, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Println(data)

	bounds := data.Bounds()

	// Loop through all the pixels on each row
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			//r, g, b, a := data.At(x, y).RGBA()
			color := data.At(x, y)
			fmt.Println(color)
			//fmt.Printf("Red: %08b, Green: %08b, Blue: %08b, Alpha: %08b\n", (r / 257), (g/ 257), (b / 257), (a / 257))		
		}
	}
}
