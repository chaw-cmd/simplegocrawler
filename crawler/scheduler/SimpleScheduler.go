package scheduler

import "simplegocrawler/crawler/engine"

type SimpleScheduler struct{
	WorkerChannel chan engine.Request
}

func (scheduler *SimpleScheduler) ConfigWorkerChannel(workerChannel chan engine.Request) {
	scheduler.WorkerChannel = workerChannel
}

func (scheduler *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		scheduler.WorkerChannel <- request
	}()
}

