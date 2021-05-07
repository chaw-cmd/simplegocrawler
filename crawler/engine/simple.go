package engine

import (
	"log"
	"simplegocrawler/crawler/fetcher"
	"time"
)

type SimpleEngine struct {}

func (se SimpleEngine) Run(seeds ... Request) {
	var requests []Request
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		parseResult, err := se.worker(request)
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

func (se SimpleEngine) worker(request Request) (ParseResult, error) {
	log.Printf("fetching: %s", request.Url)
	content, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("error fetching content from %s, error: %v", request.Url, err)
		return ParseResult{}, err
	}

	return request.ParserFunc(content), nil
}
