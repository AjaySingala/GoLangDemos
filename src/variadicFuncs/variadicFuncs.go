package main

import "fmt"

func main() {
	fmt.Println("Total is:", add(1, 2, 3))
	fmt.Println("Total is:", add(12, 13, 14, 15, 16))

	// With a slice by using "...".
	xs := []int{10, 20, 30}
	fmt.Println("Total is:", add(xs...))
}

// Println() is also a variadic function
func add(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	return total
}
