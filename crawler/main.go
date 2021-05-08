package main

import (
	"simplegocrawler/crawler/config"
	"simplegocrawler/crawler/engine"
	"simplegocrawler/crawler/persist"
	"simplegocrawler/crawler/scheduler"
	"simplegocrawler/crawler/targets/zhenai/parser"
)

func main() {
	crawlerEngine := &engine.ConcurrentEngine{
		Scheduler:       &scheduler.QueuedScheduler{}, // pointer?
		NumberOfWorkers: 3,
		PersistWorker: persist.Worker{},
	}

	crawlerEngine.Run(engine.Request{
		Url:        config.MainPage,
		ParserFunc: parser.ParseCityList,
	})
}