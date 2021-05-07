package engine

import (
	"log"
	"time"
)

type SimpleEngine struct {}

func (se *SimpleEngine) Run(seeds ... Request) {
	var requests []Request
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		parseResult, err := InputWorker(request)
		if err != nil {

			continue
		}
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("got item: %v", item)
		}
		time.Sleep(1 * time.Second) // speed limit, for local testing
	}
}

