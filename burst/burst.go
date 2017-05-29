package main

import (
	"time"
	"fmt"
)

var (
	BurstRate int 	= 3 // Burst rate of 3
)

func main() {

	burstyLimiter := make(chan time.Time, BurstRate)

	// Fill the channel to allow for an initial burst. Not necessary
	for i := 0; i < BurstRate; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			select {
			case burstyLimiter <- t: // Put t in the channel unless it's full
			default: // Necessary for not blocking when the channel is full
			}
		}
	}()

	burstyRequests := make(chan int, 5)
	go func() {
		for req := range burstyRequests {
			<-burstyLimiter
			fmt.Println("request", req, time.Now())
		}
	}()

	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}

	time.Sleep(3*time.Second)
	fmt.Println()

	for i := 1; i <= 10; i++ {
		burstyRequests <- i
	}
	time.Sleep(3*time.Second)
	close(burstyRequests)
}