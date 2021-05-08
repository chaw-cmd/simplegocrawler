package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	NumberOfWorkers int // set non-zero default??
}

type Scheduler interface {
	ReadyNotifier
	Submit(request Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (ce *ConcurrentEngine) Run(seeds ... Request) {
	// 1. init workers
	cout := make(chan ParseResult)
	ce.Scheduler.Run()
	for i := 0; i < ce.NumberOfWorkers; i++ {
		CreateInputWorker(ce.Scheduler.WorkerChan(), cout, ce.Scheduler)
	}

	// 2. send initial jobs
	for _, request := range seeds {
		if visited(request.Url) {
			// add to map for deduplication
			log.Printf("Find duplicate url: %s, skipped.", request.Url)
			continue
		}
		ce.Scheduler.Submit(request)
	}

	// 3. output items to console output
	for {
		result := <-cout
		for _, item := range result.Items {
			log.Printf("Got item: %v", item)
		}

		for _, request := range result.Requests {
			if visited(request.Url) {
				// add to map for deduplication
				log.Printf("Find duplicate url: %s, skipped.", request.Url)
				continue
			}
			ce.Scheduler.Submit(request)
		}
	}

	// use output worker(separate goroutine)
	//CreateOutputWorker(cout, ce)
	//time.Sleep(10 * time.Second)
}

// todo: bloom filter
var visitedMap = make(map[string]bool)
func visited(url string) bool {
	if visitedMap[url] {
		return true
	}

	visitedMap[url] = true
	return false
}