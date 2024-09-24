package main

import (
	"fmt"
)

func main() {
	doSomething()
	multiply := mulValues(99, 10)
	fmt.Println(multiply)

	mulmoreandmore, lengthofmul := mulMore(1, 2, 3, 4, 5)
	fmt.Printf("Mul mote = %v and it contains %v values", mulmoreandmore, lengthofmul)
}

func doSomething() {
	fmt.Println("Doing sth")
}

func mulValues(val1, val2 int) int {
	return val1 * val2
}

func mulMore(vals ...int) (int, int) {
	total := 1
	for _, v := range vals {
		total *= v
	}
	return total, len(vals)
}
