package main

import "fmt"

/*
We can handle a run-time panic with the built-in recover function.
"recover" stops the panic and returns the value that was passed to the call to panic
*/

func main() {
	/*
		In the following code, the call to recover will never happen because the call
		to panic immediately stops execution of the function.
	*/
	// panic("PANIC")
	// str := recover()
	// fmt.Println(str)

	// Instead, use this.
	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	// OR this. Not both!
	//defer tryRecover()

	panic("PANIC!!!")
}

func tryRecover() {
	str := recover()
	fmt.Println(str)
}
