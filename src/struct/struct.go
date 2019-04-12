package main

import (
	"fmt"
	"math"
)

func main() {
	// Undesired approach
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5

	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println(circleArea(cx, cy, cr))

	// Desired approach
	var c1 Circle
	fmt.Println("c1", c1)
	fmt.Println("&c1", &c1)
	fmt.Println("c1.x", c1.x)
	fmt.Println("&c1.x", &c1.x)
	fmt.Println("&c1.y", &c1.y)

	c2 := new(Circle)
	fmt.Println("c2", c2)
	fmt.Println("&c2", &c2)
	fmt.Println("*c2", *c2)
	fmt.Println("c2.x", c2.x)

	c3 := Circle{x: 0, y: 0, r: 5}
	fmt.Println("c3", c3)
	fmt.Println("&c3", &c3)
	//fmt.Println("*c3", *c3)
	fmt.Println("&c3.x", &c3.x)
	fmt.Println("&c3.y", &c3.y)
	fmt.Println("&c3.r", &c3.r)
	fmt.Println("c3.x", c3.x) // Field x
	fmt.Println("c3.x", c3.y) // Field y
	fmt.Println("c3.x", c3.r) // Field r

	// Same as c3
	c4 := Circle{0, 0, 5}
	fmt.Println("c4.x", c4.r)

	// Use funcs With structs
	// Here, changes to c4 in the function will be local to the func.
	fmt.Println(circleArea2(c4))

	// To reflect changes to c4 in the function, use address and pointers:
	fmt.Println("Before, c4.x", c4.x)
	fmt.Println(circleArea3(&c4))
	fmt.Println("After, c4.x", c4.x)
}

type Circle struct {
	x float64
	y float64
	r float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1

	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)

	return l * w
}

func circleArea(x, y, r float64) float64 {
	return math.Pi * r * r
}

func circleArea2(c Circle) float64 {
	fmt.Println("Inside circleArea2()...")
	fmt.Println("c.x", c.x)
	fmt.Println("c.r", c.r)
	return math.Pi * c.r * c.r
}

func circleArea3(c *Circle) float64 {
	fmt.Println("Inside circleArea3()...")
	fmt.Println("c.x", c.x)
	fmt.Println("c.r", c.r)

	c.x = 2
	fmt.Println("c.x", c.x)
	return math.Pi * c.r * c.r
}
