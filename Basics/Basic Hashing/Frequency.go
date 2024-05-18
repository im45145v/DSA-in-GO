package main

import (
	"fmt"
)

func freq(arr []int) {
	frequencymap := make(map[int]int)
	for _, val := range arr {
		frequencymap[val]++
	}
	fmt.Println("Element frequency")
	for k, v := range frequencymap {
		fmt.Printf("%d: %d \n", k, v)
	}
}
func main() {
	arr := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	freq(arr)
}
