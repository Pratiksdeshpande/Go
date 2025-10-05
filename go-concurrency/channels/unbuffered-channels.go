package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println("Sending message...")
		ch <- "Hi from goroutine" // blocks until received
		fmt.Println("Message sent!")
	}()

	fmt.Println("Waiting to receive...")
	msg := <-ch
	fmt.Println("Received:", msg)
}

/* Output:
Waiting to receive...
Sending message...
Message sent!
Received: Hi from goroutine


Reason for output order:

**Unbuffered Channel Execution Order Explanation**

The sequence of output messages demonstrates how goroutines and unbuffered channels work together in Go.
Here's why "Message sent!" appears before "Received: Hi from goroutine":

1. The main function starts and creates an unbuffered channel ch
2. It launches a goroutine (concurrent function)
3. The main function prints "Waiting to receive..."
4. Main function reaches msg := <-ch and blocks waiting for data on the channel
5. The scheduler switches to the goroutine, which prints "Sending message..."
6. The goroutine executes ch <- "Hi from goroutine" which transfers the message to the main function
7. At this moment, both goroutines are unblocked simultaneously:
	a. The goroutine can continue and immediately prints "Message sent!"
	b. The main function receives the value and continues
8. The main function then executes fmt.Println("Received:", msg)

The key insight is that once the channel communication happens, both goroutines are free to continue execution.
In this case, the goroutine finished its print statement before the main goroutine could execute its print statement.

This ordering isn't guaranteed - it depends on how the Go scheduler allocates CPU time between goroutines,
but this particular execution pattern is common.
*/
