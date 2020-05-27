package main

import (
	"fmt"
	"flag"
)

func main() {
	var text string
	var key string
	var option bool
	var result string

	flag.StringVar(&text, "t", "", "Text to encrypt")
	flag.StringVar(&key, "k", "", "Encryption/Decryption key")
	flag.BoolVar(&option, "d", false, "Decryption")
	flag.Parse()
	
	if option {
		for i, letter := range text {
			if int(letter) == 32 {
				result += " "
			} else if int(letter - rune(key[i % len(key)])) < 0 {
				result += string(128 + (letter - rune(key[i % len(key)])))
			} else {
				result += string(letter - rune(key[i % len(key)]))
			}
		}
	} else {
		for i, letter := range text {
			if int(letter) == 32 {
				result += " "
			} else {
				result += string(int(letter + rune(key[i % len(key)])) % 128)
			}
		}
	}
	fmt.Println(result)
}