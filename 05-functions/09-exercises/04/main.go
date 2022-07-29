package main

import "fmt"

func main() {
	b := (true && false) || (false && true) || !(false && false)
	// true && false = false
	// false && true = false
	// !(false && false) = true
	// false || false || true = true
	if b {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
