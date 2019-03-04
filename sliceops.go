package main

import "fmt"

func main() {
	var s []int

	if s == nil {
		println("emtpy--->", s)
	}
	fmt.Println(&s)
	for i := 0; i < 100; i++ {
		s = append(s, i*2+1)
	}
	fmt.Println(s)
}
