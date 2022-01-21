package main

import (
	"fmt"
	"time"
)

//	A Goroutine is a lightweight thread of execution

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	//	This is what a normal function call looks like
	f("direct")

	//	To invoke this function in a goroutine, we use `go f(s)`.
	//	This new goroutine will execute concurrently with the calling one
	go f("goroutine")

	//	You can also start a goroutine from an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	//	Wait for all goroutines to finish (in a very hacky way)
	time.Sleep(time.Second)
	fmt.Println("done")
}

//	When we run this program, we see the output of the blocking call first, then the output of the two goroutines
//	The goroutines' output may be interleaved as they are running concurrently
