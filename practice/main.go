package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	number := make([]int, 5)
	number[0] = 10
	number[1] = 7
	number[2] = 15
	number[3] = 1
	number[4] = 4
	fmt.Println(number)

	number = append(number, 19)
	fmt.Println(number)

	sort.Ints(number)
	fmt.Println(number)
}
