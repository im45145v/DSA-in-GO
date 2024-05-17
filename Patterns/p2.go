package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scan(&i)

	for row := 0; row < i+1; row++ {
		for col := 0; col < row; col++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
