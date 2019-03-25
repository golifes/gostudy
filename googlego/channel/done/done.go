package main

import (
	"fmt"
	"sync"
)

func doWorker(id int, w worker) {
	for n := range w.in {
		fmt.Printf("worker %d received %c\n", id, n)
		//go func() {
		//	done <- true
		//}()
		//done <- true
		w.done()
	}
}

type worker struct {
	in chan int
	//done chan bool
	//wg *sync.WaitGroup
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	//c := make(chan int)
	go doWorker(id, w)
	//return c
	return w
}

func chanDemo() {
	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	//wg.Add(15)

	for i, worker := range workers {
		wg.Add(1)

		worker.in <- 'a' + i
		//<-workers[i].done
	}
	for i, worker := range workers {
		wg.Add(1)

		worker.in <- 'A' + i
		//<-workers[i].done
	}

	//time.Sleep(time.Millisecond)
	//wait for all of them
	//for _, worker := range workers {
	//	<-worker.done
	//	<-worker.done
	//
	//}
	wg.Wait()
}

func main() {
	chanDemo()
}
