package main

import (
	"fmt"
)

type Base struct {
	y int
	Y int
}

func (b *Base) FuncByPoint() int {
	if b == nil {
		return 0
	}
	return b.y * b.Y
}

func (b Base) FuncByValue() int {
	return b.y * b.Y
}

type Child struct {
	Base
	x int
	X int
}

func (c *Child) FuncByPoint() int {
	if c == nil {
		return 0
	}
	return c.x * c.X
}

func main() {
	var c Child
	c.y = 2
	c.Y = 3

	fmt.Printf("%v\n", c.FuncByPoint())
	fmt.Printf("%v\n", c.Base.FuncByPoint())
	fmt.Printf("%v\n", c.FuncByValue())

	var f1 func() int
	f1 = c.FuncByPoint
	fmt.Printf("%v\n", f1())

	var f2 func(*Child) int
	f2 = (*Child).FuncByPoint()
	fmt.Printf("%v\n", f2(&c))
}
