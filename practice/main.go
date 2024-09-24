package main

import (
	"fmt"
)

func main() {
	slice := []string{"apple", "banana", "orange"}
	result := convertToMap(slice)
	fmt.Println((result))
}

func convertToMap(items []string) map[string]float64 {

	// Create a map object
	result := make(map[string]float64)

	price := 100 / float64(len(items))
	for i := range items {
		result[items[i]] = price
	}
	fmt.Println((price))

	return result
}
