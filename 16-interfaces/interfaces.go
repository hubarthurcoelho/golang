package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type form interface {
	area() float64
}

func getArea(f form) {
	fmt.Println(f.area())
}

func main() {
	rectangle := rectangle{120, 200}
	circle := circle{50}
	getArea(rectangle)
	getArea(circle)
}
