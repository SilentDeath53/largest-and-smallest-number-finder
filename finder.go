package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Enter the number of elements: ")
	fmt.Scanln(&n)

	numbers := make([]int, n)

	fmt.Printf("Enter %d numbers:\n", n)
	for i := 0; i < n; i++ {
		fmt.Printf("Number %d: ", i+1)
		fmt.Scanln(&numbers[i])
	}

	largest := numbers[0]
	smallest := numbers[0]

	for i := 1; i < n; i++ {
		if numbers[i] > largest {
			largest = numbers[i]
		}
		if numbers[i] < smallest {
			smallest = numbers[i]
		}
	}

	fmt.Printf("Largest number: %d\n", largest)
	fmt.Printf("Smallest number: %d\n", smallest)
}
