package main

import (
	"fmt"
	"time"
)

func main() {
	//for i := 0; i < 10; i++ {
	//	go func(i int) {
	//		for {
	//			fmt.Printf("Hello from goroutine %d\n", i)
	//		}
	//	}(i)
	//
	//}
	//time.Sleep(time.Minute)

	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				//runtime.Gosched()
			}
		}(i)

	}
	time.Sleep(10 * time.Millisecond)
	fmt.Println(a)
}
