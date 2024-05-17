package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scan(&i)
	for row := 0; row < i+1; row++ {
		if row==0 || row==i{
			for col := 0; col < i+1; col++ {
				fmt.Print("*")
			}
		}else{
			fmt.Print("*")
			for col := 1; col < i; col++ {
				fmt.Print(" ")
			}
			fmt.Print("*")
		}

		fmt.Println()
	}
}
