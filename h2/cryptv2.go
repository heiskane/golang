package main

import (
	"fmt"
	"io/ioutil"
	"flag"
)

var result string
var key string
var option bool
var file string

func encrypt(file string) string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Error Opening File:", err)
	} else {
		for i, letter := range bytes {
			result += string((letter + key[i % len(key)]) % 128)
		}
	}
	return result
}

func decrypt(file string) string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Error Opening File:", err)
	} else {
		for i, letter := range bytes {
			result += string(((letter - (key[i % len(key)]) % 128) + 128) % 128)
				// The % isnt actually modulus in Go but the remainder
				// To make it work correctly with netavive numbers
				// Use ( ( ( x % y ) + y ) % y )
		}
	}
	return result
}

func main() {
	
	flag.StringVar(&key, "k", "", "Encryption/Decryption key")
	flag.StringVar(&file, "f", "", "File to encrypt the content of")
	flag.BoolVar(&option, "d", false, "Decryption")
	
	flag.Parse()

	if option {
		fmt.Printf(decrypt(file))
	} else {
		// DO NOT USE "Printf" HERE OR IT CRAPS ITSELF BECAUSE OF "%" SIGNS
		fmt.Println(encrypt(file))
	}
}