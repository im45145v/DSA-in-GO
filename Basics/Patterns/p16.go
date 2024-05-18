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
	k := 'A'

	// Generate the pattern
	for i := 'A'; i <= maxChar; i++ {
		for j := 'A'; j <= i; j++ {
			fmt.Printf("%c", k)
		}
		fmt.Println()
		k += rune(1)
	}
}
