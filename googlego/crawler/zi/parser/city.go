package parser

import (
	"gostudy/googlego/crawler/engine"
	"regexp"
)

//<tr><th><a href="http://album.zhenai.com/u/1026949424" target="_blank">悦心</a></th>
//不是又尖括号  [^>]
const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<])</a>`

func ParseCity(c []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(c, -1)

	result := engine.ParseResult{}

	for _, m := range matches {
		result.Items = append(result.Items, "User ： "+string(m[2]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity,
			})
	}

	return result
}
