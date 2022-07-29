package main

import "fmt"

func main() {
	myName := "Xuan"

	SwitchOnType(myName)
}

func SwitchOnType(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("Argument is an int")
	case string:
		fmt.Println("Argument is a string")
	default:
		fmt.Println("Argument is unknown type")
	}
	// Printing: Argument is a string
}
