package main

import "fmt"

func main() {
	type bigint int64

	//var x bigint = 100
	//print(x)
	x := 1234
	var b bigint = bigint(x)
	var b2 int64 = int64(b)
	fmt.Println(b2)
}
