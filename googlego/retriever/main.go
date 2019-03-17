package main

type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster) {
	poster.Post("https://www.baidu.com", map[string]string{
		"name": "xiaohan",
	})
}

type RetrieverPoster interface {
	Poster
}

func main() {

}
