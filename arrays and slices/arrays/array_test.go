package arrays

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestArray(t *testing.T) {
	t.Run("array", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assert.Equal(t, got, want, "got %d want %d", got, want)
	})
}
