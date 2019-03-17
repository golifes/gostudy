package main

import "fmt"

func adder() func(i int) int {
	sum := 0
	return func(j int) int {
		sum += j
		return sum
	}
}

type iAdder func(i int) (int, iAdder)

func adder2(b int) iAdder {
	return func(i int) (int, iAdder) {
		return i + b, adder2(i + b)
	}
}

func main() {
	a := adder()

	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
	}
}
