/*
	Go has a special statement called select which works
	like a switch but for channels
*/
package main

import (
	"fmt"
	"time"
)

/*
	This program prints “from 1” every 2 seconds and “from 2” every 3 seconds.
	select picks the first channel that is ready and receives from it (or sends to it).
	If more than one of the channels are ready then it randomly picks which one to
	receive from. If none of the channels are ready, the statement blocks until
	one becomes available.
*/
func channelSelect() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)

				/*
					The following is to implement a timeout.
					time.After creates a channel and after the given duration
					will send the current time on it.
				*/
				// case <-time.After(time.Second):
				// 	fmt.Println("timeout")

				/*
					The default case happens immediately if none of the
					channels are ready
				*/
				// default:
				// 	fmt.Println("nothing ready")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)

	fmt.Println("select channel program ends here...")
}
