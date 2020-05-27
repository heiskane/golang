package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	b64 "encoding/base64"
)

func encode(data string) string {
	encoded := b64.StdEncoding.EncodeToString([]byte(data))
	return encoded
}

func main() {
	var file string

	flag.StringVar(&file, "f", "", "File to read and encode")
	flag.Parse()

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Error Opening file", err)
	}
	
	fmt.Println(encode(string(bytes)))
	
}