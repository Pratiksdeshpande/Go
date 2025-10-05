package main

import "fmt"

func main() {
	ch := make(chan int, 2) // buffer of 2

	ch <- 1
	fmt.Println("Sent 1")
	ch <- 2
	fmt.Println("Sent 2")

	fmt.Println("Receiving:", <-ch)
	fmt.Println("Receiving:", <-ch)
}

/*
Output:
Sent 1
Sent 2
Receiving: 1
Receiving: 2

Reason for output order:

Buffered Channels Execution Flow Explanation

This code demonstrates how buffered channels work in Go. Let me explain the output step by step:

1. ch := make(chan int, 2) creates a channel with a buffer capacity of 2 integers.
2. ch <- 1 sends the value 1 to the channel. Since the buffer has space, this operation doesn't block.
3. fmt.Println("Sent 1") executes immediately after sending the first value.
4. ch <- 2 sends the value 2 to the channel. The buffer now contains [1,2] but still isn't blocking
	because the buffer limit hasn't been exceeded.
5. fmt.Println("Sent 2") executes immediately after sending the second value.
6. fmt.Println("Receiving:", <-ch) receives the first value from the buffer (1) and prints it.
	Channels operate in FIFO (First In, First Out) order.
7. fmt.Println("Receiving:", <-ch) receives the second value from the buffer (2) and prints it.

The key insight here is that buffered channels allow sends to proceed without blocking until the buffer is full.
This is why both "Sent" messages appear before any "Receiving" messages - the sending operations completed independently
before any receiving happened.

*/
