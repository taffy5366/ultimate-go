package main

import (
	"fmt"
)

func assert(i interface{}) {
	s := i.(int)
	fmt.Println(s)
}

func main() {
	var s interface{} = 55
	assert(s)
}
