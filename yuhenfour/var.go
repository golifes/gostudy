package main

import "fmt"

func main() {
	s := "abc"
	y := "1"
	y, s = "hello", "111"

	fmt.Println(s, y)

	a := 1
	//b := 0

	//c := a/b //panic: runtime error: integer divide by zero
	c := a / 0 //division by zero
	fmt.Println(c)
}
