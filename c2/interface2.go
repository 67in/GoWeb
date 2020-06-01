package main

import (
	"fmt"
	"strconv"
)

type Human3 struct {
	name string
	age int
	phone string
}

//通过这个方法 Human 实现了 fmt.Stringer
func (h Human3) String() string {
	return "❰"+h.name+" - "+strconv.Itoa(h.age)+" years -  ✆ " +h.phone+"❱"
}

func main() {
	Bob := Human3{"Bob", 39, "000-7777-XXX"}
	fmt.Println("This Human is : ", Bob)
}
