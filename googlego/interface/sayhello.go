package main

import "fmt"

type IGreeting interface {
	sayHello()
}

func Hello(i IGreeting) {
	i.sayHello()
}

type Go struct {
}

func (g Go) sayHello() {
	fmt.Println("Hi i am GO!!!")
}

type PHP struct {
}

func (p PHP) sayHello() {
	fmt.Println("Hi i am php!!!")
}

func main() {
	golang := Go{}
	php := PHP{}

	Hello(golang)
	Hello(php)
}
