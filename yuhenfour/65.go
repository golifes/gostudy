package main

import "fmt"

type User struct {
	id   int
	name string
}

func (u *User) String() string {
	return fmt.Sprintf("%d, %v", u.id, u.name)
}

func (u *User) Print() {
	fmt.Println(u.String())
}

type Stringer interface {
	String() string
}

type Printer interface {
	String() string
	Print()
}

func main() {
	var o Printer = &User{1, "tom"}
	var s Stringer = o
	fmt.Println(s.String())
}
