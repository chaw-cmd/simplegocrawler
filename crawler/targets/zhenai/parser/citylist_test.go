package parser

import (
	"io/ioutil"
	"testing"
)

const (
	resultSize = 470
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)
	if len(result.Items) != resultSize || len(result.Requests) != resultSize {
		t.Errorf("result size not match, should be %d", resultSize)
	}
	// todo: check city names and URLs
}
