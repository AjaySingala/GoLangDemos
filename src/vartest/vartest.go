package main

import "fmt"

var global = "I am a global var"

func main() {
	var x string = "Hello there"
	fmt.Println(x)

	var y string
	y = "Hello there!"
	fmt.Println(y)

	x = "How are you doing?"
	fmt.Println(x)

	y = y + x
	fmt.Println(y)

	var a = "Hello"
	b := "Hello"
	fmt.Println(a == b)
	b = "hello"
	fmt.Println(a == b)
	/*
		b:= "some value" again will not work as b is already defined.
		b = 25 will give an error as we are assiging an int to a string.
	*/

	// Access a global var.
	fmt.Println(global)

	// Constants.
	const age int = 26
	fmt.Println(age)
	// age = 30			// error as age is a constant.
	// fmt.Println(age)

	multipleVars()
}

func multipleVars() {
	// Define multiple vars.
	var (
		a = 10
		b = 20
		c = 30
	)
	fmt.Println(a, b, c)
}

func foo() {
	fmt.Println(global)
	//fmt.Println(x)		// won't compile as x is not in scope.
}
