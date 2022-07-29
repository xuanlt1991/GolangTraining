package main

import "fmt"

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}

func foo(numbers ...int) {
	total := 0

	for _, n := range numbers {
		total += n
	}

	fmt.Println(total)
}
