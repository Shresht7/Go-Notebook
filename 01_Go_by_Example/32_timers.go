package main

import (
	"fmt"
	"time"
)

//	Timers are useful when we want to execute code at some time in the future or at an interval
//	Go has builtin timers and ticker features to make these tasks easy.

func main() {

	//	Timers represent a single event in the future.
	//	You tell the timer how long you want to wait, and it provides a channel that will be notified at the time.
	timer1 := time.NewTimer(2 * time.Second)

	//	This blocks on the timer's channel C until it sends a value indicating that the timer fired
	<-timer1.C
	fmt.Println("Timer 1 fired")

	//	If you just wanted to wait you could have used time.Sleep
	//	One reason a timer may be useful is that you can cancel the timer before it fires
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}

	//	Give the timer2 enough time to fire to show in fact that it stopped.
	time.Sleep(2 * time.Second)

}
