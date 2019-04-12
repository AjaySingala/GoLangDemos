package main

import "fmt"

func main() {
	x := 5
	fmt.Println(x)
	zero(x)
	fmt.Println(x) // still prints 5.

	fmt.Println("Memory location of x is", &x)
	zeroP(&x)      // Pass address of x
	fmt.Println(x) // prints 0.

	// Demonstrate "new()"
	demoNew()
}

func one(aPtr *int) {
	fmt.Println("aPtr =", aPtr)
	fmt.Println("&aPtr =", &aPtr)
	fmt.Println("*aPtr =", *aPtr)
	*aPtr = 1
}

func demoNew() {
	fmt.Println("In demoNew()...")
	xPtr := new(int)
	fmt.Println("xPtr =", xPtr)
	fmt.Println("&xPtr =", &xPtr)
	fmt.Println("*xPtr =", *xPtr)
	one(xPtr)
	fmt.Println("*xPtr =", *xPtr) // prints 1
}

func zero(x int) {
	x = 0
}

// receiver pointer to x.
// Pointers reference a location in memory where a value
// is stored rather than the value itself
// *int is a pointer to an integer
// &x in main() and xPtr refer to the same memory location
// *xPtr is the value of x, which is being changed.
func zeroP(xPtr *int) {
	*xPtr = 0
	fmt.Println("Memory location of x via xPtr is", xPtr)
	fmt.Println("Memory location of xPtr is (&xPtr)", &xPtr)
}
