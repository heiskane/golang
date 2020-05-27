package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main () {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}