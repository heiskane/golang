# Image Randomizer
What this does is it "randomizes" the colors based on the number that is provided by the user. Works with png and jpeg now

Example usage:

```
./imageRandomizer -f input.png -n 5 -o outfile.png
```

Compiled binaries can be found in windows/linux folder but to build it yourself you can use:

```
go build imageRandomizer.go
```

Orinally i wanted to create a steganography tool but due to time constraints it became an image randomizer.
I intend to create a separate steganograpy tool out of this in the future and maybe implement Vigenere cipher tool i have
created before.

Sources used:

https://www.socketloop.com/tutorials/golang-get-rgba-values-of-each-image-pixel
https://stackoverflow.com/questions/33186783/get-a-pixel-array-from-from-golang-image-image
https://socketloop.com/tutorials/golang-convert-integer-to-binary-octal-hexadecimal-and-back-to-integer
https://dev.to/andyhaskell/how-i-made-a-slick-personal-logo-with-go-s-standard-library-29j9
https://golang.org/pkg/image/

Images used for these examples 

http://pngimg.com/download/30524
https://www.freepngs.com/cars-pngs?pgid=jc7xa93e-4cf4c91a-5461-11e8-a9ff-063f49e9a7e4
![Messy Mario](/img/mario_messy.png)
![Messy Cars](/img/cars_messy.png)
