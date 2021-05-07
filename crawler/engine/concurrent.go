package engine

import (
	"time"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	NumberOfWorkers int // set non-zero default??
}

type Scheduler interface {
	Submit(request Request)
	ConfigWorkerChannel(workerChannel chan Request)
}

func (ce *ConcurrentEngine) Run(seeds ... Request) {
	// 1. init workers
	cin := make(chan Request)
	cout := make(chan ParseResult)
	for i := 0; i < ce.NumberOfWorkers; i++ {
		CreateInputWorkers(cin, cout)
		CreateOutputWorker(cout, ce)
	}

	// 2. send initial jobs
	ce.Scheduler.ConfigWorkerChannel(cin)
	for _, request := range seeds {
		ce.Scheduler.Submit(request)
	}

	// 3. output items to console output
	//for {
	//	result := <-cout
	//	for _, item := range result.Items {
	//		log.Printf("Got item: %v", item)
	//	}
	//
	//	for _, request := range result.Requests {
	//		ce.Scheduler.Submit(request)
	//	}
	//}

	time.Sleep(10 * time.Second)
}