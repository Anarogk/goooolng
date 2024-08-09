package main

import (
	"bytes"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
func TestCountDown(t *testing.T) {
	t.Run("CountDown", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		CountDown(buffer)
		got := buffer.String()
		want := "3\n2\n1\nGo!"
		assert.Equal(t, got, want, "got %q want %q ", got, want)
	})
}
