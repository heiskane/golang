package main

import (
	"fmt"
	"flag"
)

func main() {
	var name string
	var number int

	flag.StringVar(&name, "name", "", "Your name here")
	flag.IntVar(&number, "n", 0, "Number of greetings")
	flag.Parse()

	if name != "" {
		
		for i := 0; i <= number; i++ {
			fmt.Println("Hello", name)
		}

	} else {
		fmt.Println("Mikset kerro nimeäsi?")
	}
}