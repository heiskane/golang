# Site Grabber and some stego tests

This is my homework for the intensive Go programming course taught by Tero Karvinen http://terokarvinen.com/2020/go-programming-course-2020-w22/

The siteGrabber.go connects to a website and base64 encodes the body then outputs it.

Example usage:```./siteGrabber -u http://terokarvinen.com > outfile```

To build run ```go build siteGrabber.go```

Currently stegoTest.go only reads a png file and outputs the RGBA values but i want to try to modify the least significant bit 
and write that to a new image which would contain a hidden message.

some of the sources ive used so far:

For siteGrabber: https://stackoverflow.com/questions/38673673/access-http-response-as-string-in-go

For stegoTest: 

https://www.socketloop.com/tutorials/golang-get-rgba-values-of-each-image-pixel
 https://stackoverflow.com/questions/33186783/get-a-pixel-array-from-from-golang-image-image
 https://socketloop.com/tutorials/golang-convert-integer-to-binary-octal-hexadecimal-and-back-to-integer
 https://dev.to/andyhaskell/how-i-made-a-slick-personal-logo-with-go-s-standard-library-29j9
 https://golang.org/pkg/image/
