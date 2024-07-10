package main

import (
	"fmt"
	"math"
)

func getUniqueFactorization(x int) []int {
	var factors []int

	if x%2 == 0 {
		factors = append(factors, 2)
		for x%2 == 0 {
			x /= 2
		}
	}

	sqrtX := int(math.Sqrt(float64(x)))
	for i := 3; i <= sqrtX; i += 2 {
		if x%i == 0 {
			factors = append(factors, i)
			for x%i == 0 {
				x /= i
			}
			sqrtX = int(math.Sqrt(float64(x)))
		}
	}

	if x > 2 {
		factors = append(factors, x)
	}

	return factors
}

func main() {
	var x int

	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&x)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Printf("Unique prime factorization for %d: ", x)
	p := getUniqueFactorization(x)
	for _, factor := range p {
		fmt.Printf("%d ", factor)
	}
}
