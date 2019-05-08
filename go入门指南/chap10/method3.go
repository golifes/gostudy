package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {

	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamePoint struct {
	p    Point //如果这里是匿名结构体,下面则可以使用 n.Abs()直接访问
	name string
}

func main() {
	n := &NamePoint{Point{3, 4}, "xiaohan"}
	fmt.Println(n.p.Abs())
	fmt.Println(n.p.Abs())
}
