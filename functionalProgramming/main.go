package main

import (
	"LearnGo/functionalProgramming/fib"
	"bufio"
	"fmt"
	"io"
	"strings"
)

// func fibonacci() intGenerator {
// 	a, b := 0, 1
// 	return func() int {
// 		a, b = b, a+b
// 		return a
// 	}
// }

type intGenerator func() int

func (generator intGenerator) Read(p []byte) (n int, err error) {
	next := generator()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	var f intGenerator = fib.Fibonacci()
	printFileContents(f)

	// fmt.Println(f())
	// fmt.Println(f())
	// fmt.Println(f())
	// fmt.Println(f())
	// fmt.Println(f())
	// fmt.Println(f())
	// fmt.Println(f())
	// fmt.Println(f())
	// fmt.Println(f())

}
