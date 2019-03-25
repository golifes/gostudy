package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//fetch, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")

	contents, err := ioutil.ReadFile("citylist_test_data.html")

	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n", contents)

	result := ParseCityList(contents)

	const resultSize = 470

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests ; but had %d ", resultSize, len(result.Requests))
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d  ; but had %d ", resultSize, len(result.Items))
	}
}
