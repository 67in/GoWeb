package main

import "fmt"

type Rectangel struct {
	width, height float64
}

func area(r Rectangel) float64  {
	return  r.width* r.height
}

func main()  {
	r1 := Rectangel{12, 2}
	r2 := Rectangel{9, 4}
	fmt.Println("area of r1 is: ",area(r1))
	fmt.Println("area of r2 is: ",area(r2))

}
