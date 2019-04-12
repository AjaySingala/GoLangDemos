package main

import "fmt"

func main() {
	fmt.Println("Factorial of 2")
	fmt.Println(factorial(2))
	fmt.Println("Factorial of 10")
	fmt.Println(factorial(10))
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}

	y := x * factorial(x-1)
	fmt.Println(y)
	return y
}
