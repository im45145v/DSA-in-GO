package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Enter a number (1-26):")
	fmt.Scanln(&n)

	if n < 1 || n > 26 {
		fmt.Println("Invalid input. Please enter a number between 1 and 26.")
		return
	}

	// Calculate the maximum character based on the input
	maxChar := 'A' + rune(n) - 1

	// Generate the pattern
	for i := maxChar; i >= 'A'; i-- {
		for j := i; j <= maxChar; j++ {
			fmt.Printf("%c", j)
		}
		fmt.Println()
	}
}
