package main

import "fmt"

func main() {
	myName := "Xuan"
	switch myName {
	case "Xuan":
		fmt.Println("Hello Xuan")
		fallthrough
	case "Hai":
		fmt.Println("Hello Hai")
	default:
		fmt.Println("Hello Unknown")
	}
	// Printing: Hello Xuan \n Hello Hai
}
