package main

import (
	"fmt"
	"regexp"
)

const text = "My email ccmouse@gmail.com.cn"

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+\.[a-zA-Z0-9]+)`)

	//if err != nil{
	//	panic(err)
	//}

	match := re.FindAllStringSubmatch(text, -1)

	fmt.Println(match)
}
