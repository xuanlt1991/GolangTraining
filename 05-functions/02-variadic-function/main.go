package main

import "fmt"

func main() {
	ave := average(30, 40, 50, 60, 70, 80)
	fmt.Println(ave)
}

func average(numbers ...float64) float64 {
	fmt.Printf("%T\n", numbers)
	var total float64

	for _, n := range numbers {
		total += n
	}

	return total / float64(len(numbers))
}
