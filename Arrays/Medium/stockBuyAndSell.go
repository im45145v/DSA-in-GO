package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	profit := 0
	minPrice := prices[0]
	maxPrice := prices[0]

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
			maxPrice = 0
		} else {
			maxPrice = price
			if maxPrice-minPrice > profit {
				profit = maxPrice - minPrice
			}
		}
	}

	return profit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Max profit:", maxProfit(prices))
}
