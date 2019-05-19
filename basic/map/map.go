package main

import "fmt"

func lengthOfNonRepearingSubStr(s string) int {
	// lastOcurred := make(map[byte]int)
	lastOcurred := make(map[string]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		// lastIndex, ok := lastOcurred[ch]
		if lastIndex, ok := lastOcurred[string(ch)]; ok && lastIndex >= start {
			start = lastIndex + 1
		}

		fmt.Println("i, ch = ", i, string(ch))
		fmt.Println("start = ", start)
		fmt.Println("lastOcurred map = ", lastOcurred)
		fmt.Println("currentLength, maxLength = ", i-start+1, maxLength)
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOcurred[string(ch)] = i
		// lastOcurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println("Ans = ", lengthOfNonRepearingSubStr("abcabcbb"))
	fmt.Println("Ans = ", lengthOfNonRepearingSubStr("pwwkew"))
	fmt.Println("Ans = ", lengthOfNonRepearingSubStr("bbbbb"))
}
