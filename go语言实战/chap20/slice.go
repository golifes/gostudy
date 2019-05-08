package main

import "fmt"

func main() {
	//var wh [3]struct{}
	//for i := range wh {
	//	defer func() {
	//		fmt.Println(i)
	//	}()
	//}
	//
	//var a intc
	//
	//defer fmt.Println(a)
	//
	//a = 4
	r := f()
	fmt.Println("r--->", r)

	f2 := f2()
	fmt.Println(f2)

}

func f2() (r int) {
	t := 5

	defer func() {
		t = t + 5
		fmt.Println("inner ttt", t)
	}()
	fmt.Println("ttt", t)
	return t
}
func f() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)

	return 1
}
