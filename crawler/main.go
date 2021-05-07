package main

import (
	"simplegocrawler/crawler/config"
	"simplegocrawler/crawler/engine"
	"simplegocrawler/crawler/scheduler"
	"simplegocrawler/crawler/targets/zhenai/parser"
)

func main() {
	crawlerEngine := &engine.ConcurrentEngine{
		Scheduler:       &scheduler.SimpleScheduler{}, // pointer?
		NumberOfWorkers: 3,
	}

	crawlerEngine.Run(engine.Request{
		Url:        config.MainPage,
		ParserFunc: parser.ParseCityList,
	})
}