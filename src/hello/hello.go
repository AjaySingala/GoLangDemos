package main

import "fmt"

func main() {
	fmt.Println("Hello there!")

	fmt.Println("Hello " + "there!")
	fmt.Println("Hello", "there!")
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1 + 1 =", 1.0+1.52)

	fmt.Println(len("Hello there!"))
	fmt.Println("Ajay Singala"[0])
	fmt.Println("Hello there!"[1])

	boolTest()
}

func boolTest() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
