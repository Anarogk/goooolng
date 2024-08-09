package main

import "fmt"

// Add takes two floats and returns their sum
func Add(a, b float64) float64 {
	return a + b
}
func main() {
	fmt.Println(Add(10, 10))
}
