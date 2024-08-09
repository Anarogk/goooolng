package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	endWord = "Go!"
	count   = 3
)

type Sleeper interface {
	Sleep()
}

func CountDown(std_out io.Writer) {
	for i := count; i > 0; i-- {
		fmt.Fprintln(std_out, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(std_out, "Go!")
}
func main() {
	CountDown(os.Stdout)
}
