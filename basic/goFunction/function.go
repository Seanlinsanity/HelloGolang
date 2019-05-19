package main

import "fmt"
import "reflect"
import "runtime"
import "math"

func divide(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	pointer := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(pointer).Name()
	fmt.Printf("calling function %s with args"+"(%d, %d)\n", opName, a, b)

	return op(a, b)
}

func myPower(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sum(numbers ...int) int {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

func swap(a, b *int) {
	*b, *a = *a, *b
}

func swapB(a, b int) (int, int) {
	return b, a
}

func main() {
	q, r := divide(10, 3)
	fmt.Println(q, r)
	fmt.Println(apply(myPower, 2, 4))
	fmt.Println(apply(
		func(a, b int) int {
			return a * b
		}, 10, 20,
	))
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
	c, d := 5, 6
	c, d = swapB(c, d)
	fmt.Println(c, d)
}
