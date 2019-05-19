package main

import (
	"fmt"
	"math"
	"strconv"
)

var (
	name   = "Stephen Curry"
	number = 30
	isCool = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "what's up"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "Wow"
	b = 5
	fmt.Println(a, b, c, s)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const (
		stephen = "stephen"
		a, b    = 3, 4
	)
	//GLOBALCONSTANT global string here
	const GLOBALCONSTANT string = "globalConstantHere"
	var c int
	c = a + b*b
	fmt.Println(a, b, c, stephen)

}

func enums() {
	const (
		swift = iota
		javascript
		objectiveC
		golang
		python
	)

	const (
		b = 1 << (iota * 10)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(swift, python, javascript, objectiveC, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func basicLoop() {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func convertToBin(num int) string {
	result := ""
	if num == 0 {
		return "0"
	}
	for ; num > 0; num /= 2 {
		lsb := num % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func forever() {
	for {
		fmt.Println("run loop...")
	}
}

func main() {
	fmt.Println("Sean write in Go here...")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(name, number, isCool)
	triangle()
	consts()
	enums()
	basicLoop()
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
		convertToBin(859213041),
		convertToBin(0),
	)
	// forever()

}
