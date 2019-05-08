package main

func main() {
	//	var i int
	//	for {
	//		println(i)
	//		i++
	//
	//		if i > 3 {
	//			goto Break
	//		}
	//	}
	//
	//Break:
	//	println("break")

	continue_goto()

}

func continue_goto() {
L1:
	for x := 0; x < 3; x++ {
	L2:
		for y := 0; y < 5; y++ {
			if y > 2 {
				continue L2
			}

			if x > 1 {
				break L1
			}

			print(x, ":", y, " ")
		}

		println()
	}
}
