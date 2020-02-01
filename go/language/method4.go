package main

import (
	"fmt"
)

type myInt struct {
	owner string
	value int
}

func (a myInt) Owner(suffix string) string { //golang不支持默认参数
	return a.owner + suffix
}

func (a *myInt) SetOwner(owner string) {
	if a == nil {
		fmt.Println("set owner to nil point is invalid")
		return
	}
	a.owner = owner
}

func (a myInt) SetOwner2(owner string) {
	a.owner = owner //golang函数参数安值传递，所以这个方法实际只是修改临时变量的owner
	fmt.Println("%v", a)
}

func SetOwner3(a *myInt, owner string) {
	if a == nil {
		fmt.Println("set owner to nil point is invalid")
		return
	}
	a.owner = owner
}

func main() {
	var k = myInt{"kitman", 3}

	fmt.Print(k.value, " ", k.Owner("aa"), "\n") //3 kitmanaa

	k.SetOwner("ak")
	fmt.Print(k.value, " ", k.Owner("bb"), "\n") //3 akbb

	k.SetOwner2("sss")
	fmt.Printf("%v\n", k)
	fmt.Print(k.value, " ", k.Owner("bb"), "\n") //3 akbb

	SetOwner3(&k, "sss")
	fmt.Printf("%v\n", k)
	fmt.Print(k.value, " ", k.Owner("bb"), "\n") //3 sssbb

	var k2 *myInt = nil
	k2.SetOwner("aa")
}
