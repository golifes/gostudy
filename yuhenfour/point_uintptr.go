package main

import "fmt"

func main() {
	d := struct {
		s string
		x int
	}{"abc", 10}

	fmt.Println(d)
}
