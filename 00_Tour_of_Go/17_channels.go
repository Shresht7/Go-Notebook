package main

import "fmt"

//	Channels are typed conduit through which you can send and receive values with the channel operator <-
//	```
//	ch <- y	//	Send y to channel ch
//	x := <- ch	//	Receive from ch, and assign to x
//	```
//	The data flows in direction of the arrow

//	Like maps and slices, channels must be created before use
//	`ch := make(chan int)`

//	By default, send and receive block until the other side is ready.
//	This allows goroutines to synchronize without explicit locks or condition variables.

//	This example code sums the numbers in a slice, distributing the work between two goroutines
//	Once both goroutines have completed their computation, it calculates the final result.

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //	Send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c //	Received from c

	fmt.Println(x, y)

	//	Channels can be buffered by providing the buffer length as the second argument to make
	buf := make(chan int, 2)
	//	Sends to a buffer block only when the buffer is full
	//	Receives block when the buffer is empty.
	buf <- 1
	buf <- 2
	buf <- 3
	fmt.Println(<-buf)
	fmt.Println(<-buf)
}
