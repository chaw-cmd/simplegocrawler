package parser

import (
	"regexp"
	"simplegocrawler/crawler/engine"
)

const cityListRegex = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	reg := regexp.MustCompile(cityListRegex)
	matched := reg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matched {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: engine.NilParser, // todo: next level parse func
		})
	}
	return result
}