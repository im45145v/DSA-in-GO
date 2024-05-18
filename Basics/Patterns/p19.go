package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scan(&i)

	for row := 0; row < i+1; row++ {

		for col := 0; col < (i-row)+1; col++ {
			fmt.Print("*")
		}
		for col := 0; col < 2*(row); col++ {
			fmt.Print(" ")
		}
		for col := 0; col < (i-row)+1; col++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for row := i; row > 0; row-- {

		for col := 0; col < (i-row)+1; col++ {
			fmt.Print("*")
		}
		for col := 0; col < 2*(row); col++ {
			fmt.Print(" ")
		}
		for col := 0; col < (i-row)+1; col++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
