package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkNum(0))
}
func checkNum(a int) string {
	var result string

	if a < 0 {
		result = "Less than 0"
	} else if a == 0 {
		result = "Is Zero"
	} else {
		result = "Greater than 0"
	}
	return result
}
