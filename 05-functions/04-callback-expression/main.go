package main

import "fmt"

func main() {

	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})

	fmt.Println(xs) // [2 3 4]
}

func filter(numbers []int, callback func(int) bool) []int {
	fmt.Printf("Type of callback: %T\n", callback)
	xs := []int{}

	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}

	return xs
}
