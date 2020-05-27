package main

import (
	"fmt"
)

func dogYears(i int) (int, int) {
	return i * 7, i * 6
}

func main() {
	fmt.Println(dogYears(10))
}