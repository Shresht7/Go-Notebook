package main

import (
	"fmt"
	"time"
)

//	Timeouts are import for programs that connect to external resources
//	or that otherwise need to bound execution time. Implementing timeouts in Go is easy and elegant
//	thanks to channels and select

func main() {

	c1 := make(chan string, 1)
	//	For this example, assume our call takes 2s to complete.
	//	Note that this channel is buffered, so sending in the goroutine is non-blocking
	//	This is a common pattern to prevent goroutine leaks in case the channel is never read
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	//	Here's the select implementing the timeout
	//	res := c1 awaits the result and <-time.After awaits a value to be sent after the timeout of 1s.
	//	Since the select proceeds with the first receive that's ready, we'll take the timeout case if the operation takes more than the allowed 1s.
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

}

//	Running this program shows that the first operation timed out and the second succeeded.
