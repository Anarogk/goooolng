package arrays

import "fmt"

func Sum(numbers [5]int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	var n int = Sum(a)
	fmt.Println(n)
}
