package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync/atomic"
// 	"time"
// )

//	In addition to mutexes, another option is to use the builtin sync features of goroutines
//	and channels to achieve the same result. This channel-base approach aligns with Go's idea of sharing memory by
//	communiation and having each piece of data owned by exactly 1 goroutine

//	In this example, our state will be owned by a single goroutine. This will guarantee
//	that that data is never corrupted with concurrent access
//	In order t oread or write that state, other goroutines will send messages to the owning goroutine
//	and receive the corresponding replies.
//	These readOp and writeOp structs encapsulate those requests and a way for owning the goroutine to respond
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	//	As before, we'll count how many operations we perform
	var readOps uint64
	var writeOps uint64

	//	These channels will be used by other goroutines to issue read and write requests
	reads := make(chan readOp)
	writes := make(chan writeOp)

	//	Here is the goroutine that owns the state, which is a map.
	//	This goroutine repeatedly selects on the reads and write channels,
	//	responding to requests as they arrive.
	//	A response is executed by first performing the requested operation and then
	//	sending a value on the response channel resp to indicate success.
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	//	This starts 100 goroutines to issue reads to the state-owning goroutine via read's channel
	//	TODO: Revisit Goroutines and complete these examples
}
