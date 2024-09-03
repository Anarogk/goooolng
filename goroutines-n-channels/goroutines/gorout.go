package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 100; i++ {
			fmt.Println("number from first goroutine", i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 0; i < 100; i++ {
			fmt.Println("number from Second goroutine", i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		}
	}()

	wg.Wait()
	fmt.Println("Done")

}
