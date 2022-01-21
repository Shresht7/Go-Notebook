package main

import (
	"fmt"
	"time"
)

//	We can use channels to synchronize executions across goroutines

//	This is the function we'll run in a goroutine. The done channel will be used
//	to notify another goroutine that this function's work is done
func worker(done chan bool) {
	fmt.Println("working ...")
	time.Sleep(time.Second)
	fmt.Println("done")

	//	Send a value to notify that we're done
	done <- true
}

func main() {

	done := make(chan bool, 1)
	//	Start the worker goroutine and give it the channel to notify on
	go worker(done)

	//	Block until we receive a notification form the worker on the channel
	<-done
}

//	If you remove the <- done line from this program, the program would exit before the worker even started.
