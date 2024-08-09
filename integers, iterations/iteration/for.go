package iteration

import "fmt"

func ForLoop(a string, n int) string {
	var s string
	for i := 0; i < n; i++ {
		s += a
	}
	return s
}
func main() {
	fmt.Println(ForLoop("a", 5))
}
