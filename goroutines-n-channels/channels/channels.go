package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1) // creating a channel with buffer size 1

	go func() {
		sum := 0

		for i := 0; i < 100; i++ {
			fmt.Println("number from first function", i)
			sum += i
		}
		ch <- sum // sending sum to channel
	}()

	// blocking process of receiving sum from channel it does not terminate until the sum is received
	output := <-ch // receiving sum from channel
	fmt.Println("sum from first goroutine", output)
}
