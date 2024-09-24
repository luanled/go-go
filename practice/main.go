package main

import (
	"fmt"
	"math/rand"
)

func main() {

	dow := rand.Intn(7) + 1
	fmt.Println("Day", dow)

	var res string
	switch dow {
	case 1:
		res = "Sunday the king plays"
	case 2:
		res = "It's Monday!"
	case 3:
		res = "It's Tuesday!"
	case 4:
		res = "It's Wednesday!"
	case 5:
		res = "It's Thursday!"
	case 6:
		res = "It's Friday!"
	case 7:
		res = "It's Saturday!"
	}
	fmt.Println(res)
}
