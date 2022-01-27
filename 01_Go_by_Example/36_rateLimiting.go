package main

import (
	"fmt"
	"time"
)

//	Rate limiting is an import mechanism for controlling resource utilization and
//	maintaining quality of service. Go supports rate limiting with goroutines, channels
//	and tickers

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	//	this limiter channel will receive a value every 200 milliseconds. This is our regulator
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		//	By blocking on a receive from the limiter channel before serving each requests
		//	we limit ourselves to 1 request every 200 milliseconds
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	//	We may want to allow short bursts for request in our rate limiting scheme while preserving the overall limit
	//	We can accomplish this by buffering our channel. This burstyLimiter channel will allow bursts of up to 3 events
	burstyLimiter := make(chan time.Time, 3)

	//	Fill up the channel to allow bursting
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	//	Every 200ms we'll try to add a new value to bursty limiter, up to its limit of 3
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	//	Now simulate 5 more incoming requests. The first 3 of these will benefit from the burst capability.
	burstyRequests := make(chan int, 5)
	for i := 5; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

//	Running our program we see the first batch of requests handled once every ~200ms as desired.
//	For the second batch of requests, we serve the first 3 immediately because of the burstable rate limiting,
//	then serve the remaining 2 with 200ms delay each.
