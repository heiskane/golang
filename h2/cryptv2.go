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
			if letter - key[i % len(key)] < 0 {
				result += string(128 + letter - (key[i % len(key)]))
			} else {
				result += string((letter - (key[i % len(key)])) % 128)
			}
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