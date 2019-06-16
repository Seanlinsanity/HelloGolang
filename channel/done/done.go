package main

import (
	"fmt"
	"sync"
)

type worker struct {
	receiver chan int
	// done     chan bool
	done func()
}

func doWork(id int, c chan int, worker worker) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		// done <- true
		worker.done()
	}
}

func createWorker(id int, waitGroup *sync.WaitGroup) worker {
	worker := worker{
		receiver: make(chan int),
		// done:     make(chan bool),
		done: func() {
			waitGroup.Done()
		},
	}
	go doWork(id, worker.receiver, worker)
	return worker
}

func channelDemo() {
	var waitGroup sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &waitGroup)
	}

	waitGroup.Add(20)

	for i, worker := range workers {
		worker.receiver <- 'a' + i
	}

	// for _, worker := range workers {
	// 	<-worker.done
	// }

	for i, worker := range workers {
		worker.receiver <- 'A' + i
	}

	// for _, worker := range workers {
	// 	<-worker.done
	// }

	waitGroup.Wait()
}

func main() {
	channelDemo()
}
