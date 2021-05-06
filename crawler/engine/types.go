package engine

type Request struct {
	Url string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request // new requests parsed from current page
	Items []interface{} // items parsed from current page
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
