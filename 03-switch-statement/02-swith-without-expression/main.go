package main

import "fmt"

func main() {
	myName := "James"

	switch {
	case myName == "Smith", myName == "James":
		fmt.Println("Hello Smith or James")
	case myName == "Xuan":
		fmt.Println("Hello Xuan")
	default:
		fmt.Println("Hello Unknown")
	}
	// Printing: Hello Smith or James
}
