package main

import (
	"fmt"
)

func testDefer() {
	defer fmt.Println("hello")
	defer fmt.Println("hello v2")
	defer fmt.Println("hello v3")
	fmt.Println("aaaaaa")
	return
	fmt.Println("bbbbbb")
}

func testDefer2() {
	var i int = 0
	defer fmt.Printf("defer i=%d\n", i)
	i = 100
	fmt.Printf("i=%d\n", i)
}
func main() {
	testDefer()
	testDefer2()
}
