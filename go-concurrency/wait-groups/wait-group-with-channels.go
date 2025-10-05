package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	results := make(chan int, len(nums))
	var wg sync.WaitGroup

	for _, n := range nums {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			results <- x * x
		}(n)
	}

	wg.Wait()
	close(results)

	for r := range results {
		fmt.Println("Result:", r)
	}
}
