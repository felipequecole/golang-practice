package main

import "fmt"

func main() {
	fmt.Println(average(2, 3))
}

func average(xi ...float64) float64 {
	entriesCount := float64(len(xi))
	var totalSum float64
	for _, value := range xi {
		totalSum += float64(value)
	}

	return totalSum / entriesCount
}
