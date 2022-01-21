package main

import (
	"fmt"
	"time"
)

//	Go's select lets you wait on multiple channel operations
//	Combining goroutines and channels with select is a powerful feature of Go

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	//	Here we're using two channels,
	//	each channel will receive a value after some amount of time, to simulate a blocking operation in concurrent goroutines

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	//	We'll use select to await both of these values simultaneously, printing each one as it arrives
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received c1:", msg1)
		case msg2 := <-c2:
			fmt.Println("received c2:", msg2)
		}
	}
}

//	We receive the values "one" and "two" as expected.
//	Note that the total execution time is only ~2 seconds since both 1 and 2 seconds Sleeps execute concurrently.
