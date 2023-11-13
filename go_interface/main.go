package main

import "fmt"

type Rectangle struct {
	height int
	width  int
}

type Paralegram struct {
	main   int
	height int
}

type Shape interface {
	area() float64
}

func (r Rectangle) area() float64 {
	return float64(r.height) * float64(r.width)
}

func (p Paralegram) area() float64 {
	return float64(p.height) * float64(p.main)
}

func calculate(s Shape)float64 {
	return s.area()
}

func main() {

	rect := Rectangle{3, 4}
	par := Paralegram{3, 4}

	fmt.Println(calculate(rect))
	fmt.Println(calculate(par))
}