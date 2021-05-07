package engine

import (
	"log"
	"simplegocrawler/crawler/fetcher"
)

func InputWorker(request Request) (ParseResult, error) {
	log.Printf("fetching: %s", request.Url)
	content, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("error fetching content from %s, error: %v", request.Url, err)
		return ParseResult{}, err
	}

	return request.ParserFunc(content), nil
}

func CreateInputWorkers(cin chan Request, cout chan ParseResult) {
	go func() {
		for {
			request := <-cin
			result, err := InputWorker(request)
			if err != nil {
				continue
			}
			cout <- result
		}
	}()
}


func OutputWorker(result ParseResult, ce *ConcurrentEngine) {
	for _, item := range result.Items {
		log.Printf("Got item: %v", item)
	}

	for _, request := range result.Requests {
		ce.Scheduler.Submit(request)
	}
}

func CreateOutputWorker(cout chan ParseResult, ce *ConcurrentEngine) {
	go func() {
		for {
			OutputWorker(<-cout, ce)
		}
	}()
}
