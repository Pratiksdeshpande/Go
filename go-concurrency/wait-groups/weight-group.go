package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// main() exits immediately, workers may not finish
	//concurrencyWithoutWaitGroup()

	concurrencyWithWaitGroup()
}

func concurrencyWithoutWaitGroup() {
	for i := 1; i <= 3; i++ {
		go worker(i)
	}
	fmt.Println("All workers started")
}

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func concurrencyWithWaitGroup() {
	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1) // we are starting a goroutine
		go workerWithWaitGroup(i, &wg)
	}

	wg.Wait() // wait until all goroutines are finished
	fmt.Println("All workers completed")
}

func workerWithWaitGroup(id int, wg *sync.WaitGroup) {
	defer wg.Done() // mark this goroutine as done at the end
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d done\n", id)
}
