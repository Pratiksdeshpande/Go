package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		print("Hello from a goroutine!")
		wg.Done()
	}()

	wg.Wait()
	println("Main function completed.")
}

func print(str string) {
	fmt.Println(str)
	time.Sleep(2 * time.Second)
	fmt.Println("Again " + str)
	time.Sleep(2 * time.Second)
}
