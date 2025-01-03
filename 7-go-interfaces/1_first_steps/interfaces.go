package main

import (
	"fmt"
	"math"
)

type figures2D interface {
	area() float64
}

type rectangle struct {
	length float64
	width  float64
}

type square struct {
	base float64
}

func (s square) area() float64 {
	area := math.Pow(s.base, 2)
	return area
}

func (r rectangle) area() float64 {
	area := r.length * r.width
	return area
}

func calculateArea(figure figures2D) {
	fmt.Println("Area:", figure.area())
}

func main() {
	square1 := square{base: 15.5}
	rectangle1 := rectangle{length: 30, width: 20}
	calculateArea(square1)
	calculateArea(rectangle1)
}
