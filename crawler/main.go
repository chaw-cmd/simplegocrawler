package main

import (
	"simplegocrawler/crawler/config"
	"simplegocrawler/crawler/engine"
	"simplegocrawler/crawler/targets/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        config.MainPage,
		ParserFunc: parser.ParseCityList,
	})
}