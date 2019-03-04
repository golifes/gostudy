package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes 你好！"

	bytes := []byte(s)

	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c --%d", ch, size)
	}
}
