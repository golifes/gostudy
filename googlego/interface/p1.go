package main

import (
	"fmt"
)

type Person struct {
	Name string
}
type a int

func (a) Study(lang string) {
	fmt.Println("study language:", lang)
}
func main() {
	var p a
	p.Study("Go")
}
