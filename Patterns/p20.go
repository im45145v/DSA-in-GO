package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scan(&i)

	for row := 1; row < i; row++ {
		for col := 0; col < row; col++ {
			fmt.Print("*")
		}
		for col := 0; col < 2*(i-row); col++ {
			fmt.Print(" ")
		}
		for col := 0; col < row; col++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for row := i; row > 0; row-- {
		for col := 0; col < row; col++ {
			fmt.Print("*")
		}
		for col := 0; col < 2*(i-row); col++ {
			fmt.Print(" ")
		}
		for col := 0; col < row; col++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
