package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Shape interface {
	area() float64
}

func main() {
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}
	fmt.Println("Area of Circle c:", totalArea(&c))
	fmt.Println("Area of Rectangle r:", totalArea(&r))
	fmt.Println("Total Area of r and c:", totalArea(&c, &r))
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}

	return area
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)

	return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1

	return math.Sqrt(a*a + b*b)
}
