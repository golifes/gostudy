package main

func lengthOfNonRepeatingSubStr(s string) int {

	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {

		last, ok := lastOccurred[ch]

		if ok && last >= start {
			start = lastOccurred[ch] + 1

		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}

		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	ret := lengthOfNonRepeatingSubStr("abcab")
	println(ret)
}
