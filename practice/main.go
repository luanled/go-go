package main

import (
	"fmt"
)

func main() {
	ipadPro := Device{"iPadOS", 799}
	fmt.Println(ipadPro)
	fmt.Printf("%+v\n", ipadPro)
}

// Struct test
type Device struct {
	OS    string
	Price int
}
