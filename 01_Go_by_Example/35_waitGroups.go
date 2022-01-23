package main

import (
	"fmt"
	"sync"
	"time"
)

//	To wait for multiple goroutines to finish, we can use a wait group

//	This is the function we'll run in every goroutine
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) //	Sleep to simulate expensive task
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	//	This wait group is used to wait for all goroutines launched here to finish.
	//	Note: if a WaitGroup is explicitly passed into functions, it should be done by pointers
	var wg sync.WaitGroup

	//	Launch several goroutines and increment the WaitGroup counter for each
	for i := 1; i <= 5; i++ {
		wg.Add(i)

		//	Avoid reusing of the same i value in each goroutine closure.
		i := i

		//	Wrap the worker call in a closure that makes sure to tell the WaitGroup that
		//	this worker is done. This way the worker itself does not have to be aware of the concurrency primitives involved in its execution.
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	//	Block until the WaitGroup counter goes back to 0; all they worker notified they're done
	wg.Wait()
}

//	The order of workers staring up and finishing is likely to be different for each invocation
