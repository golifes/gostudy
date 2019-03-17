package main

import (
	"fmt"
	"time"
)

func worker(c chan int) {

	for {
		n := <-c
		fmt.Println(n)

	}
}

func chanDemo() {
	c := make(chan int)

	go worker(c)

	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
}
func main() {
	bufferedChannel()
	//chanDemo()
}
