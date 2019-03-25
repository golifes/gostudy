package main

import (
	"fmt"
	"time"
)

func generator() chan int {
	out := make(chan int)
	i := 0
	go func() {
		for {
			time.Sleep(time.Millisecond)
			out <- i
			i++
		}
	}()

	return out
}

func main() {
	var c1, c2 = generator(), generator()
	for {

		select {
		case n := <-c1:
			fmt.Println("received from c1", n)
		case n := <-c2:
			fmt.Println("received from c2", n)
			//default:
			//	fmt.Println("No value received")

		}
	}
}
