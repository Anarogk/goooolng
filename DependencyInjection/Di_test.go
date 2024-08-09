package main

import (
	"bytes"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	t.Run("DItest", func(t *testing.T) {
		buf := bytes.Buffer{}
		// The Buffer type from the bytes package implements the Writer interface,
		// because it has the method Write(p []byte) (n int, err error).
		Greet(&buf, "Chris")

		got := buf.String()
		want := `Hello, Chris`

		assert.Equal(t, got, want, "got %q want %q", got, want)
	})
}
