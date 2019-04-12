package main

import "fmt"

// Function within a function is "closure"
func main() {
	fmt.Println("Closure test...")
	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(10, 20))

	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

	// 	One way to use closure is by writing a function which
	// returns another function which – when called – can
	// generate a sequence of numbers.
	fmt.Println("Even numbers...")
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // prints 0
	fmt.Println(nextEven()) // prints 2
	fmt.Println(nextEven()) // prints 4

}

// makeEvenGenerator returns a function which generates
// even numbers. Each time it's called it adds 2 to the local
// i variable which – unlike normal local variables –
// persists between calls.
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
