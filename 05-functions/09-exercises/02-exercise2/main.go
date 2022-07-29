package main

import "fmt"

func main() {
	x := 1
	y := 2
	half := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}

	fmt.Println(half(x))
	fmt.Println(half(y))
}
