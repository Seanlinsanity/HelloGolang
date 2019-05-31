package main

import (
	"LearnGo/functionalProgramming/fib"
	"bufio"
	"fmt"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occured")
}

func tryDeferWithLoop() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("panic and stop")
		}
	}
}

func writeFile(filename string) {
	// file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	file, err := os.Create(filename)
	if err != nil {
		if pathError, success := err.(*os.PathError); !success {
			panic(err)
		} else {
			fmt.Printf("operation: %s, path: %s, error: %s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	writeFile("fibonacci.txt")
	tryDeferWithLoop()
}
