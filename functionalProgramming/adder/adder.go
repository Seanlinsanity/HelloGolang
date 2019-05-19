package main

import "fmt"

func adder() func(int) int {
	//sum: free variable
	sum := 0
	return func(v int) int {
		//v: local variable
		sum += v
		return sum
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		// fmt.Println(a(i))
		fmt.Printf("i = %d, a(i) = %d\n", i, a(i))
	}

}
