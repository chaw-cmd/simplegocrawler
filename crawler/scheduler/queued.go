package scheduler

import "simplegocrawler/crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request // worker's request chan and worker is 1:1 mapping
}

func (qs *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (qs *QueuedScheduler) Submit(request engine.Request) {
	qs.requestChan <- request
}

func (qs *QueuedScheduler) WorkerReady(w chan engine.Request) {
	qs.workerChan <- w
}

func (qs *QueuedScheduler) Run() {
	qs.requestChan = make(chan engine.Request)
	qs.workerChan = make(chan chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var readyRequest engine.Request
			var readyWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				readyRequest = requestQ[0]
				readyWorker = workerQ[0]
			}

			select {
			case r := <-qs.requestChan:
				requestQ = append(requestQ, r)
			case w := <-qs.workerChan:
				workerQ = append(workerQ, w)
			case readyWorker <- readyRequest:
				// if queues are empty, readyWorker will be nil
				// then this case won't be reached
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}

