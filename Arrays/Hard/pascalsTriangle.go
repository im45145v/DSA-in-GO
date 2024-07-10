package main

import "fmt"

func generate(numRows int) [][]int {
	finalNums := [][]int{}
	if numRows <= 0 {
		return finalNums
	}

	finalNums = append(finalNums, []int{1})
	for i := 1; i < numRows; i++ {
		newRow := make([]int, i+1)
		newRow[0], newRow[i] = 1, 1
		for j := 1; j < i; j++ {
			newRow[j] = finalNums[i-1][j-1] + finalNums[i-1][j]
		}
		finalNums = append(finalNums, newRow)
	}
	return finalNums
}

func main() {
	numRows := 5
	fmt.Printf("Pascal's Triangle with %d rows:\n", numRows)
	result := generate(numRows)
	for _, row := range result {
		fmt.Println(row)
	}
}
