package main

import (
	"fmt"
)

type Square struct {
	side uint
}

type Rectangle struct {
	width  uint
	height uint
}

type Shape interface {
	GetPerimeter() uint
	GetArea() uint
}

func main() {
	shapes := []Shape{&Square{side: 2}, &Rectangle{height: 5, width: 3}}
	var totalArea uint
	var totalPerimeter uint

	for _, shape := range shapes {
		totalArea += shape.GetArea()
		totalPerimeter += shape.GetPerimeter()
	}

	fmt.Println("Total area of Square and Rectangle is", totalArea)
	fmt.Println("Total Perimeter of Square and Rectangle is", totalPerimeter)

}

func (s *Square) GetPerimeter() uint {
	var perimeter uint = s.side * 4
	fmt.Println("perimeter of square:", perimeter)

	return perimeter
}

func (s *Square) GetArea() uint {
	var area uint = s.side * s.side
	fmt.Println("Area of shape:", area)

	return area
}

func (r *Rectangle) GetPerimeter() uint {
	var perimeter uint = (r.width + r.height) * 2
	fmt.Println("perimeter of rectangle:", perimeter)

	return perimeter
}

func (r *Rectangle) GetArea() uint {
	var area uint = r.width * r.height
	fmt.Println("Area of rectangle:", area)

	return area
}
