package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scan(&i)
	p := 0

	for row := 0; row < i+1; row++ {
		for col := 0; col < row; col++ {
			p++
			fmt.Print(p)
		}
		for col := 0; col < 2*(i-row); col++ {
			fmt.Print(" ")
		}
		for col := 0; col < row; col++ {
			fmt.Print(p)
			p--
		}

		p = 0
		fmt.Println()
	}
}
