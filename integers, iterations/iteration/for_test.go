package iteration

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

// Test for for function in main package in integers, iterations
func TestForloop(t *testing.T) {
	t.Run("for", func(t *testing.T) {
		got := ForLoop("a", 5)
		want := "aaaaa"
		assert.Equal(t, got, want, "got %s want %s", got, want)
	})
}

// Benchmarking
func BenchmarkForLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ForLoop("a", 3)
	}
}
