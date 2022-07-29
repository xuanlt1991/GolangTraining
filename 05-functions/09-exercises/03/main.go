package main

import "fmt"

func main() {
	max := findMax([]int{4, 1, 68, 90, 23, 11, 91}...)

	fmt.Println(max)
}

func findMax(numbers ...int) int {
	max := 0

	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	return max
}
