package main

import (
	"fmt"
	"time"
)

type myTime struct {
	t time.Time
}

func (t myTime) first3Chars() string {
	return t.t.String()[0:4]
}

func main() {
	m := myTime{time.Now()}

	fmt.Println("Full time now:", m.t.String())

	fmt.Println("First 4 chars : ", m.first3Chars())

}
