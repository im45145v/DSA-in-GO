package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scan(&i)
	p := 0

	for row := 0; row < i+1; row++ {
		for col := 0; col < i-row; col++ {
			p++
			fmt.Print(p)
		}
		p = 0
		fmt.Println()
	}
}
