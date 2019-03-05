package main

import "fmt"

func main() {
	s := "abcd汉字"
	bs := []rune(s)
	bs[1] = 'B'
	print(string(bs))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%c,\n", s[i])
	}

	for index, r := range s {
		fmt.Printf(" %c | %d ", r, index)
	}

}
