package main

import "fmt"

func main() {
	fmt.Println(factorial(4)) // 24
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}

// Step 1: 4 => return 4 * 6 = 24
// Step 2: 3 => return 3 * 2 = 6
// Step 3: 2 => return 2 * 1 = 2
// Step 4: 1 => return 1 * 1 = 0
// Step 5: 0 => return 1
