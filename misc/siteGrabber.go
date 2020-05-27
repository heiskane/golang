package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"flag"
	b64 "encoding/base64"
)

// https://gobyexample.com/base64-encoding
func encode(data string) string {
	encoded := b64.StdEncoding.EncodeToString([]byte(data))
	return encoded
}

func main () {
	var url string

	flag.StringVar(&url, "u", "http://example.com/", "Url of the site to grab")
	flag.Parse()

	resp, err := http.Get(flag)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	// https://stackoverflow.com/questions/38673673/access-http-response-as-string-in-go
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { 
		fmt.Println(err)
	}
	
	bodyString := string(body)
	fmt.Printf(encode(bodyString))
}