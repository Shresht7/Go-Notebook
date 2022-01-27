package main

import (
	"fmt"
	"sync"
)

//	For more complex state we can use a mutex to safely access data across multiple goroutines

//	Container holds a map of counters; since we want to update it concurrently from multiple goroutines,
//	we add a Mutex to synchronize access. note that mutexes must not be copied, so it this struct is passed around,
//	it should be done by a pointer.
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

//	Increment function. Lock the mutex before accessing the counters. Unlock it at the end using a defer statement
func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	//	Note that the zero value of a mutex is usable as-is, so no initialization is required here
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	//	Wait group
	var wg sync.WaitGroup

	//	This function increments a named counter in a loop
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	//	Run several goroutines concurrently, note that they all access the same Container, and two of them access the same counter
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	//	Wait for the goroutines to finish
	wg.Wait()

	//	Running the program shows the counters updated as expected.
	fmt.Println(c.counters)
}
