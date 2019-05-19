package main

import (
	"LearnGo/queue"
	"fmt"
)

func main() {
	q := queue.Queue{1}
	q.Push(11)
	q.Push(30)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	q.Push("Sean")
	fmt.Println(q.Pop())
}
