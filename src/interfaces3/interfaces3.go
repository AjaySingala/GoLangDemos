// Base implementation for interface.
// Add caching to square and rectangle.

package main

import (
	"fmt"
)

type Shape interface {
	GetPerimeter() uint
	GetArea() uint
}

func main() {
	shapes := []Shape{
		&Square{Cache{cache: make(map[string]uint)}, 2},
		&Rectangle{Cache{cache: make(map[string]uint)}, 3, 5}}
	var totalArea uint
	var totalPerimeter uint

	// First time, data is not in cache.
	fmt.Println("First time, no cache!")
	for _, shape := range shapes {
		totalArea += shape.GetArea()
		totalPerimeter += shape.GetPerimeter()
	}

	fmt.Println("Total area of Square and Rectangle is", totalArea)
	fmt.Println("Total Perimeter of Square and Rectangle is", totalPerimeter)

	// Again. This time, takes values from the cache.
	fmt.Println("This time, takes values from the cache")
	for _, shape := range shapes {
		totalArea += shape.GetArea()
		totalPerimeter += shape.GetPerimeter()
	}

	fmt.Println("Total area of Square and Rectangle is", totalArea)
	fmt.Println("Total Perimeter of Square and Rectangle is", totalPerimeter)

}

/*
	Define a Cache struct that can store the value of previous computations.
	When a value is retrieved from the case, it also prints to the screen "cache hit",
	and when the value is not in the case, it prints "cache miss" and returns -1
*/
type Cache struct {
	cache map[string]uint
}

func (c *Cache) GetValue(name string) int {
	value, ok := c.cache[name]
	if ok {
		fmt.Println("Cache hit")
		return int(value)
	} else {
		fmt.Println("Cache missing")
		return -1
	}
}

func (c *Cache) SetValue(name string, value uint) {
	c.cache[name] = value
}

/*
	Embed this cache in the Square and Rectangle shapes.
	Note that the implementation of GetPerimeter() and GetArea() now checks the
	cache first and computes the value only if it is not in the cache.
*/
type Square struct {
	Cache
	side uint
}

func (s *Square) GetPerimeter() uint {
	value := s.GetValue("perimeter")
	if value == -1 {
		value = int(s.side * 4)
		s.SetValue("perimeter", uint(value))
	}
	fmt.Println("perimeter of square:", value)

	return uint(value)
}

func (s *Square) GetArea() uint {
	value := s.GetValue("area")
	if value == -1 {
		value = int(s.side * s.side)
		s.SetValue("area", uint(value))
	}
	fmt.Println("Area of square:", value)

	return uint(value)
}

type Rectangle struct {
	Cache
	width  uint
	height uint
}

func (r *Rectangle) GetPerimeter() uint {
	value := r.GetValue("perimeter")
	if value == -1 {
		value = int((r.width + r.height) * 2)
		r.SetValue("perimeter", uint(value))
	}
	fmt.Println("perimeter of rectangle:", value)

	return uint(value)
}

func (r *Rectangle) GetArea() uint {
	value := r.GetValue("area")
	if value == -1 {
		value = int(r.width * r.height)
		r.SetValue("area", uint(value))
	}
	fmt.Println("area of rectangle:", value)

	return uint(value)
}
