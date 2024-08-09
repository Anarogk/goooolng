package main

import (
	_ "fmt"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func AreaTest(t *testing.T) {
	t.Run("Area for ", func(t *testing.T) {
		circle := Circle{10.00}
		got := circle.Area()
		want := 314.1592653589793
		assert.Equal(t, got, want, "got: %q and want: %q", got, want)
	})
}
