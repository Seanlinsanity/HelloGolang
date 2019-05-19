package main

import "fmt"
import "unicode/utf8"

func main() {
	s := "Golang練習!" //UTF-8
	fmt.Printf("%X\n", []byte(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, char := range s {
		//char is a rune
		fmt.Printf("(i = %d, char = %X)", i, char)
	}
	fmt.Println()

	fmt.Println("Rune count: ", utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		char, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c\n", char)
	}
	fmt.Println()

	for i, char := range []rune(s) {
		fmt.Printf("(%d, %c) ", i, char)
	}
	fmt.Println()

}
