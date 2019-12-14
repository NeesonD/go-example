package main

import (
	"fmt"
	"math"
)

func main() {
	r := rect2{width: 10, height: 20}
	c := circle{radius: 18}

	measure(r)
	measure(c)
}

type geometry interface {
	area() float64
	perim() float64
}

type rect2 struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect2) area() float64 {
	return r.width * r.height
}
func (r rect2) perim() float64 {
	return 2*r.height + 2*r.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(geometry2 geometry) {
	fmt.Println(geometry2)
	fmt.Println(geometry2.area())
	fmt.Println(geometry2.perim())
}
