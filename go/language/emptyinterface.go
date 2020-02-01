package main

import (
	"fmt"
)

func main() {
	s := "Hello World"
	i := 55
	strt := struct {
		name string
	}{name: "Naveen R"}
	describe(s)
	describe(i)
	describe(strt)
}

func describe(i interface{}) {
	fmt.Printf("Type = %T,value = %v\n", i, i)
}
