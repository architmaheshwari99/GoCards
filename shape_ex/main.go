package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5*t.base + 0.5*t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	t := triangle{12, 12}
	s := square{5}

	printArea(t)
	printArea(s)
}
