package main

import "fmt"

//	A sender can close a channel to indicate that no more values will be sent.
//	Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression
//	v, ok <- ch
//	if ok if false, there are no more values to receive and the channel is closed

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)

	go fibonacci(cap(c), c)

	//	The loop receives values from the channel repeatedly until it is closed
	for i := range c {
		fmt.Println(i)
	}
}

//	NOTE: only sender should close a channel, never the receiver.
//	Sending on a closed channel will cause a panic.

//	Another NOTE: Channels aren't like files. You don't usually need to close them.
//	Closing them is only necessary when the receiver must be told there are no more values comming
//	such as to terminate a range loop
