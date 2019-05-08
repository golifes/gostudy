package main

import "fmt"

func main() {
	b := []int{1, 2, 3}

	bb := make([]*int, 0, 3)
	for k, v := range b {
		println("k=>v", k, v)
		d := v
		bb = append(bb, &d)
		//bb = append(bb, &b[k])
		fmt.Println("k-v--->", &v, bb, v)
	}
	fmt.Println(bb)
	for _, v := range bb {
		fmt.Println(*v)

	}

}
