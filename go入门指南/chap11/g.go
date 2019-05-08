package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)

	if <-ch == "London" {
		fmt.Println(<-ch)
	}
	//time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}

func getData(ch chan string) (a chan map[string]interface{}) {
	var input string
	// time.Sleep(1e9)
	for {
		a["ok"] = <-ch
		fmt.Printf("%s ", input)
		return a
	}
	//return a
}
