package main

import "fmt"

//	Closing a channel indicates that no more values will be sent on it
//	This can be useful to communicate completion to the channel's receivers

func main() {
	//	We use jobs channel to communicate work to be done from the main() goroutine to a worker goroutine
	//	When we have no more jobs for the worker we'll close the jobs channel
	jobs := make(chan int, 5)
	done := make(chan bool)

	//	The worker goroutine repeatedly receives jobs.
	//	more value will be false if jobs has been closed and all values in the channel have already been received
	//	We use this to notify on done when we've worked all our jobs
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	//	This sends 3 jobs to the worker over the jobs channel then closes it
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)

	//	We await the worker using the synchronization approach
	<-done
}
