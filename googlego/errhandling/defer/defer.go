package main

import "fmt"

func main() {
	tryDefer()
}

func writeFile() {

}

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)

	fmt.Println(3)
}

type UserError interface {
	error
	Message() string
}
