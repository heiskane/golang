package main

import (
	"fmt"
)

func main() {
	var planets = []string {"Merkurius", "Venus", "Maa", "Mars"}
		for i, planet := range planets {
			fmt.Println(i, planet)
		}
}