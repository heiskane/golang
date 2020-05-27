package main

import (
	"fmt"
	"flag"
)

func main () {
	var number int

	flag.IntVar(&number, "n", 0, "Amount of times to loop")
	flag.Parse()

	num := 0
	for i := 0; i < number; i++ {
		num += 1
		if num % 2 == 0 {
			fmt.Println("\u2713 Hellö Wörld!")
		} else {
			fmt.Println(num)
		}
	}
}