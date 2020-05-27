package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	b64 "encoding/base64"
)

func encode(data string) string {
	encoded := b64.StdEncoding.EncodeToString([]byte(data))
	return encoded
}

func main () {
	resp, err := http.Get("http://example.com/")
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
	fmt.Println(encode(bodyString))
}