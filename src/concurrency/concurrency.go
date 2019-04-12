package main

import (
	"fmt"
)

func main() {

	// Basic concurrency test in Go
	//concurrencyTest()

	/*
		This is to demonstrate "channels". Channels provide a way for two goroutines
		to communicate with one another and synchronize their execution.
	*/
	//testChannels()

	/*
		This is to demonstrate select with channels in Go.
		Select is like "switch", but for channels.
		The method is in the file select.go
	*/
	channelSelect()
}

func testChannels() {
	var c chan string = make(chan string)

	// These are in another file named channels.go
	go pinger(c)

	// // Call ponger() along with pinger() and see what happens.
	// go ponger(c)
	// // The program will now take turns printing “ping” and “pong”.

	go printer(c)

	var input string
	fmt.Scanln(&input)

	fmt.Println("channel program ends here...")
}

func concurrencyTest() {
	/*
		With a goroutine we return immediately to the next line and
		don't wait for the function to complete.
	*/
	go foo(0)

	// // Run the goroutine 10 times.
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Calling foo #", i)
	// 	go foo(i)
	// }

	/*
		With a goroutine we return immediately to the next line and
		don't wait for the function to complete.
		A "goroutine" runs concurrently. If we do not have the following "Scanln"
		the program immediately goes to println without printing the numbers
		in foo()
	*/
	var input string
	fmt.Scanln(&input)

	fmt.Println("goroutine program ends here...")
}

func foo(n int) {
	for i := 0; i < 10; i++ {
		// It seems to run the goroutines in order rather than simultaneously
		fmt.Println(n, ":", i)

		/*
			Add some delay to the function using time.Sleep() and rand.Intn().
			It prints out the numbers from 0 to 10, waiting between 0 and 250 ms
			after each one. The goroutines should now run simultaneously.
		*/
		// amt := time.Duration(rand.Intn(250))
		// time.Sleep(time.Millisecond * amt)
	}
}
