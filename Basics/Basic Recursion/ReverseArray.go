package main

import (
	"fmt"
)

func reverse(arr []int, start int, end int) {
	if start >= end {
		return
	}
	arr[start], arr[end] = arr[end], arr[start]
	reverse(arr, start+1, end-1)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Original array:", arr)
	reverse(arr, 0, len(arr)-1)
	fmt.Println("Reversed array:", arr)
}
