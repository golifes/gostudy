package main

import "fmt"

type P struct {
	Age int
}

func (p *P) growUp() {
	p.Age++
}

func main() {
	m := map[string]P{
		"zhangsan": P{Age: 20},
	}

	p := m["zhangsan"]
	p.Age = 23
	p.growUp()
	m["zhangsan"] = p
	fmt.Println(m)
}
