package main

import (
	"fmt"
	"io/ioutil"
	"flag"
)

func encrypt(file string, key string) (string, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	var result string
	for i, letter := range bytes {
		result += string((letter + key[i % len(key)]) % 128)
	}
	return result, nil
}

func decrypt(file string, key string) (string, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	var result string
	for i, letter := range bytes {
		result += string(((letter - (key[i % len(key)]) % 128) + 128) % 128)
			// To make this work correctly with netavive numbers
			// Use ( ( ( x % y ) + y ) % y )
	}
	return result, nil
}

func main() {
	
	var key string
	var isDecrypt bool
	var file string

	flag.StringVar(&key, "k", "", "Encryption/Decryption key")
	flag.StringVar(&file, "f", "", "File to encrypt the content of")
	flag.BoolVar(&isDecrypt, "d", false, "Decryption")
	
	flag.Parse()

	var err error
	var result string
	if isDecrypt {
		result, err = decrypt(file, key)
	} else {
		result, err = encrypt(file, key)
	}
	if err != nil {
		fmt.Println("Failed:", err)
	}
	fmt.Printf("%v\n", result)
}