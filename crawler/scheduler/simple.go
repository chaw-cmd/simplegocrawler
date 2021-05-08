package scheduler

import "simplegocrawler/crawler/engine"

type SimpleScheduler struct{
	WorkerChannel chan engine.Request
}

func (scheduler *SimpleScheduler) WorkerChan() chan engine.Request {
	return scheduler.WorkerChannel
}

func (scheduler *SimpleScheduler) WorkerReady(w chan engine.Request) {
	panic("implement me")
}

func (scheduler *SimpleScheduler) Run() {
	scheduler.WorkerChannel = make(chan engine.Request)
}

func (scheduler *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		scheduler.WorkerChannel <- request
	}()
}

