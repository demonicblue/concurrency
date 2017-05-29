package main

import (
	// "time"
	"fmt"
	"time"
)

type Job struct {
	Payload int
}

func (j Job) Process() error {
	defer Wg.Done()
	time.Sleep(100*time.Millisecond)
	fmt.Printf("Processed %d\n", j.Payload)
	return nil//fmt.Errorf("%d", j.Payload)
}

type Worker struct {
	WorkerPool  chan chan Job
	JobChannel  chan Job
	quit    	chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool)}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	fmt.Printf("worker: Starting\n")
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				// we have received a work request.
				if err := job.Process(); err != nil {
					fmt.Errorf("Error processing: %s", err)
				}

			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

