package main

import (
	"gostudy/googlego/crawler/engine"
	"gostudy/googlego/crawler/scheduler"
	"gostudy/googlego/crawler/zi/parser"
)

func main() {

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 50,
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
