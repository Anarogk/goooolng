package main

import (
	"fmt"
	"testing"
)

// Test for add function in main package in integers, iterations

func TestAdd(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		got := Add(10.0, 10.0)
		want := 20.00
		assertCorrectMessage(t, got, want)
	})
}

// gets testing.T and compares got and want
func assertCorrectMessage(t *testing.T, got, want float64) {
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}

// Testable Examples are compiled whenever tests are executed.
// If ever your code changes so that the example is no longer valid, your build will fail.
func ExampleAdd() {
	sum := Add(10.0, 10.0)
	fmt.Println(sum)
	// Output: 20

}

// adding this comment inside ExampleAdd means the example will also be executed
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(10.0, 10.0)
	}
}
