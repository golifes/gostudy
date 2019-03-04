package main

import "fmt"

func main() {
	m2 := make(map[string]string)

	if len(m2) == 0 {
		fmt.Println("m2 is nil", m2)
	}

	var m3 map[string]string
	if m3 == nil {
		fmt.Println("m3 is nil", m3)
	}

}
