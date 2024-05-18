package main

import "fmt"

func maxFrequency(nums []int) int {
	// Create a map to count the frequency of each element
	frequencyMap := make(map[int]int)
	for _, val := range nums {
		frequencyMap[val]++
	}

	// Find the maximum frequency
	maxFreq := 0
	for _, freq := range frequencyMap {
		if freq > maxFreq {
			maxFreq = freq
		}
	}
	return maxFreq
}

func main() {
	// Example usage
	nums := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	result := maxFrequency(nums)
	fmt.Println("The maximum frequency is:", result)
}
