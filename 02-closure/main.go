package main

import "fmt"

func main() {
	increment1 := wrapper()
	increment2 := wrapper()
	fmt.Println(increment1())
	fmt.Println(increment1())
	fmt.Println(increment2())

}

func wrapper() func() int {
	x := 0

	return func() int {
		x++
		return x
	}
}
