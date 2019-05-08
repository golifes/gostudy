package main

import (
	"fmt"
	"unsafe"
)

type File2 struct {
	fd   int
	name string
}

func NewFile(fd int, name string) *File2 {
	if fd < 0 {
		return nil
	}
	fmt.Println(unsafe.Sizeof(&File2{}))
	return &File2{fd, name}
}

func main() {
	f := NewFile(10, "go入门指南/chap10/vcard.go")
	fmt.Println(f)
}
