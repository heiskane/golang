package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 17; i++ {
		if i % 2 == 0 {
			fmt.Println(" --- --- --- --- --- --- --- ---")
		} else {
			fmt.Println("|   |   |   |   |   |   |   |   |")
		}
	}
}