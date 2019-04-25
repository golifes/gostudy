package main

import (
	"fmt"
)

type Per struct {
	Name string
}

func (p Per) value() {
	fmt.Println(" p is value method accept ")
}

func (p *Per) Point() {
	fmt.Println(" p is point method accept")

}
func main() {
	//值
	var v Per
	fmt.Printf("%v %s", v, v) //{} {}
	v.value()
	v.Point()

	//point
	var p *Per //正确姿势  p := new(Per)
	// panic 因为 var只是定义了p，没有初始化 因此 panic: runtime error: invalid memory address or nil pointer dereference

	p.value()
	p.Point()

	//
	Per{}.value()
	Per{}.Point() //cannot call pointer method on Per literal

	//因此这里我个人认为涉及到用正确姿势初始化的问题,初始化方式不对，会引起panic

	//疑问 Per{}这样初始化到底跟 var v Per有什么不同
	fmt.Printf("%v", Per{}) //{}

	/**

	 */

}
