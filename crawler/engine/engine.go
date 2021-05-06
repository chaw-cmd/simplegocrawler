package engine

import (
	"log"
	"simplegocrawler/crawler/fetcher"
	"time"
)

func Run(seeds ... Request) {
	var requests []Request
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		log.Printf("fetching: %s", request.Url)
		content, err := fetcher.Fetch(request.Url)
		if err != nil {
			log.Printf("error fetching content from %s, error: %v", request.Url, err)
		}

		parseResult := request.ParserFunc(content)
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("got item: %v", item)
		}
		time.Sleep(3 * time.Second) // speed limit, for local testing
	}

}
