package main

import (
	"fmt"
)

func main() {
	var result int
	for i := 0; i <= 1000; i++ {
		result += i
		fmt.Println(result)
	}
}