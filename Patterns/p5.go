package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scan(&i)

	for row := i; row > 0; row-- {
		for col := 0; col < row; col++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
