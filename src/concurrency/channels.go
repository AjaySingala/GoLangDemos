package main

import (
	"fmt"
	"time"
)

/*
	This program will print “ping” forever (hit enter to stop it). A channel type is
	represented with the keyword chan followed by the type of the things that are
	passed on the channel (in this case we are passing strings).
	The <- (left arrow) operator is used to send and receive messages on the channel.
	c <- "ping" means send "ping". msg := <- c means receive a message and store it
	in msg. The fmt line could also have been written like this:
	fmt.Println(<-c) in which case 	we could remove the previous line.
	Using a channel like this synchronizes the two goroutines.
	When pinger attempts to send a message on the channel it will wait until printer
	is ready to receive the message. (this is known as blocking)
*/
func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

// Another method. Call this along with pinger() and see what happens.
func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}
