package main

import (
	"fmt"
)

func tryRecover() {
	defer handlePanic()
	// panic(errors.New("my new error"))

	// b := 0
	// a := 5 / b
	// fmt.Println(a)

	panic(123)
}

func handlePanic() {
	rec := recover()
	if err, ok := rec.(error); ok {
		fmt.Println("Error occured: ", err)
	} else {
		panic(fmt.Sprintf("can not handle panic: %v", rec))
	}
}

func main() {
	tryRecover()
}
