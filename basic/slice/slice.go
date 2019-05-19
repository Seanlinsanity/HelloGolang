package main

import "fmt"

func playWithSlice() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] = ", arr[2:6])
	s1 := arr[2:]
	fmt.Println("s1 = ", s1)
	s2 := arr[:6]
	fmt.Println("s2 = ", s2)
	s3 := arr[:]
	fmt.Println("arr[:] = ", s3)

	updateSlice(s1)
	fmt.Println("update s1 =", s1)
	fmt.Println("after update s1 arr = ", arr)
	updateSlice(s2)
	fmt.Println("update s2 =", s2)
	fmt.Println("after update s2 arr = ", arr)

	fmt.Println("Reslice")
	s3 = s3[:5]
	fmt.Println("relice s3 = ", s3)
	s3 = s3[2:]
	fmt.Println("relice s3 = ", s3)
}

func updateSlice(s []int) {
	s[0] = 100
}

func extendSlice() {
	fmt.Println("Extending Slice")
	arr2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr2)
	s4 := arr2[2:6]
	s5 := s4[3:5]
	fmt.Printf("s4 = %v, length(s4)=%d, capacity(s4)=%d\n", s4, len(s4), cap(s4))
	fmt.Printf("s5 = %v, length(s5)=%d, capacity(s5)=%d", s5, len(s5), cap(s5))

	s6 := append(s5, 100)
	s7 := append(s6, 101)
	s8 := append(s7, 102)
	fmt.Println("\narr2 = ", arr2)
	fmt.Println("s6 = ", s6)
	//s7 and s8 no longer view arr2
	fmt.Println("s7 = ", s7)
	fmt.Println("s8 = ", s8)
}

func printSlice(s []int) {
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))
}

func initSlice() {
	var s []int
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, i*2+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	fmt.Println("s2 = ", s2)
	printSlice(s2)
	s3 := make([]int, 10, 32)
	fmt.Println("s3 = ", s3)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1)
	fmt.Println("s2 =", s2)
	printSlice(s2)

	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	fmt.Println("s2 =", s2)
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println("front, s2 =", front, s2)
	printSlice(s2)

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println("tail, s2 =", tail, s2)
	printSlice(s2)
}

func main() {
	playWithSlice()
	extendSlice()
	initSlice()
}
