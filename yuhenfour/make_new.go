package main

import "fmt"

func main() {
	a := []int{0, 0, 0}
	a[1] = 10

	fmt.Println(a)

	b := make([]int, 3)
	b[1] = 10

	fmt.Println(b)
	c := new([]int)
	c = &b
	fmt.Println(c)

}
