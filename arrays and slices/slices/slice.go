package slices

import "fmt"

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	a := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println(a)
}
