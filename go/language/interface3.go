package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //匿名字段，默认Student就包含了Human的所有字段
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

//Human对象实现SayHi方法
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Human对象事项Sing方法
func (h *Human) Sing(lyrics string) {
	fmt.Println("La La,la la,la la la la la...", lyrics)
}

//Human对象实现Guzzle方法
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

//Employee重载Human的Sayhi方法
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s,I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

//Student 实现BorryMony方法
func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount
}

//Employee实现SpendSalary方法
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount
}

//定义interface
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

func main() {

}
