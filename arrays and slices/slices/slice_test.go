package slices

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	t.Run("Slice", func(t *testing.T) {
		mySlice := []int{1, 2, 3, 4, 5}
		got := Sum(mySlice)
		want := 15
		assert.Equal(t, got, want, "got %d want %d", got, want)
	})

	t.Run("Slice", func(t *testing.T) {
		mySlice := []int{3, 7, 9, 11, 13}
		got := Sum(mySlice)
		want := 15
		assert.Equal(t, got, want, "got %d want %d", got, want)
	})
}
