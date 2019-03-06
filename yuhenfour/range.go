package main

func main() {
	a := [3]int{0, 1, 2}

	print(a[0], " ", &a[0], " ", a[1], " ", &a[1], " ", &a)
}
