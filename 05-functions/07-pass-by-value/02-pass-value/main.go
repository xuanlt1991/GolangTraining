package main

import "fmt"

func main() {

	age := 44
	changeMe(age)
	fmt.Println(age) // 44
}

func changeMe(z int) {
	fmt.Println(z) // 44
	z = 24
	fmt.Println(z) // 24
}
