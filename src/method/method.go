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

/*
	Method is a special type of function.
	In between the keyword func and the name of the function we've added a “receiver”.
	The receiver is like a Structs and Interfaces parameter – it has a name and a type –
	but by creating the function in this way it allows us to call the function
	using the "." operator. No need for the "&" operator anymore.
*/

func main() {
	// Circle.
	c := Circle{0, 0, 5}
	fmt.Println("Circle's Area:", c.area())

	// Rectangle.
	r := Rectangle{0, 0, 10, 10}
	fmt.Println("Rectangle's Area:", r.area())

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
