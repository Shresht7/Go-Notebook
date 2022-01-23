package main

import (
	"fmt"
	"time"
)

//	Timers are for when you want to do something once
//	Tickers are for when you want to do something repeatedly at regular intervals

func main() {

	//	Tickers use a similar mechanism to timers, a channel that is sent values.
	//	Here we'll use the select builtin on the channel to await the values as they arrive every 500ms
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	//	Tickers can be stopped like timers
	//	Once a ticker is stooped it wont receive any more values on its channels
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
