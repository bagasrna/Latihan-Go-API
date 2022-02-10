package main

import (
	"fmt"
)

func main() {
	scores := []int{10, 5, 8, 9, 6}
	total := sum(scores)
	fmt.Println(total)
}

func sum(numbers []int) int {
	var total int
	for i := 0; i < len(numbers); i++ {
		total = total + numbers[i]
	}
	return total
}
