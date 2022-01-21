package main

import (
	"fmt"
	"sync"
	"time"
)

//	If we just want to make sure that only one goroutine can access a variable at a time.
//	we need to use a data structure called a mutex (mutual exclusion)
//	Go's standard library provides mutual exclusion with sync.Mutex and its two methods: Lock and Unlock.

//	SafeCounter is safe to use concurrently
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

//	We can define a block of code to be executed in mutex by surrounding it with a call to Lock and Unlock

//	Inc increments the counter for the given key
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	//	Lock so only one goroutine can access it at a time
	c.v[key]++
	c.mu.Unlock()
}

//	A value returns the current value of the counter for the given key
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock() //	We can use the defer keyword to ensure the mutex will be unlocked at the end
	return c.v[key]
}

func main() {
	c := SafeCounter{
		v: make(map[string]int),
	}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
