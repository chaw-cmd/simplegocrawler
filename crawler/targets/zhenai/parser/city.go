package parser

import (
	"regexp"
	"simplegocrawler/crawler/engine"
)

const cityRegex = `<a href="(.*album\.zhenai\.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

var (
	cityUrlRe = regexp.MustCompile(
		`href="(.*www\.zhenai\.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult {
	reg := regexp.MustCompile(cityRegex)
	matched := reg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matched {
		name := string(m[2])
		result.Items = append(result.Items, name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(contents []byte) engine.ParseResult {
				return ParseProfile(contents, name)
			}, // todo: next level parse func
		})
	}

	matches := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				ParserFunc: ParseCity,
			})
	}
	return result
}
