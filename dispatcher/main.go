package main

import (
	// "time"
	// "fmt"
	"sync"
	"runtime"
	"fmt"
)

var (
	MaxWorker       = runtime.NumCPU()// 3 //os.Getenv("MAX_WORKERS")
	MaxQueue        = 20 //os.Getenv("MAX_QUEUE")
	MaxLength int64 = 2048
	NumJobs int = 100
)

var JobQueue chan Job
var Wg sync.WaitGroup

func main() {
	JobQueue = make(chan Job, MaxQueue)
	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()

	func() {
		for i := 0; i < NumJobs; i++ {
			// fmt.Printf("Queueing job #%d\n",i)
			Wg.Add(1)
			JobQueue <- Job{Payload:i}
		}
	}()

	//time.Sleep(100*time.Second)
	Wg.Wait()
	fmt.Println("All done!")

}