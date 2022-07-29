package main

import "fmt"

func main() {
	m := make([]string, 1, 25) // Make map, slice, channel
	fmt.Println(m)             // []
	changeMe(m)
	fmt.Println(m) // [Xuan]
}

func changeMe(s []string) {
	s[0] = "Xuan"
	fmt.Println(s) // [Xuan]
}
