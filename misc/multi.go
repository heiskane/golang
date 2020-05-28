package main

import (
	"fmt"
	"time"
)

func testingMore(ender chan string) {
	fmt.Println("Func3")
	time.Sleep(1*time.Second)
	fmt.Println("Func3 Done")
	ender <- "Bye World"
}

func testingStuff(ender chan string) {
	fmt.Println("Secondary func")
	time.Sleep(3*time.Second)
	fmt.Println("Secondary Done")
	ender <- "Hello World"
}

func main() {
	ch := make(chan string)
	fmt.Println("Main func")

	//go testingStuff(ch)
	//go testingStuff(ch)
	//go testingStuff(ch)

	//go testingMore(ch)
	
	for i := 0; i < 4; i++ {
		go testingStuff(ch)
		go testingMore(ch)
	}

	for i := 0; i < 4; i++ {
		test := <- ch
		fmt.Println(test)
	}

	//test1 := <- ch
	//test2 := <- ch
	//test3 := <- ch
	//test4 := <- ch

	//fmt.Println("Test1:", test1)
	//fmt.Println("Test2:", test2)
	//fmt.Println("Test3:", test3)
	//fmt.Println("Test4:", test4)
	fmt.Println("Primary Done")
}