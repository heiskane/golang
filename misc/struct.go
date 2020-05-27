package main

import (
	"fmt"
)

type test struct {
	x string
	y string
	z int
}

func main() {
	var p test
	p.x = "Peruna"
	p.y = "Tomaatti"
	fmt.Println(p)
}