package main

import (
	"fmt"
)

func main() {
	var str string = "This is Go!"
	fmt.Println(str)
	fmt.Printf("str type is %T\n", str)

	var num int = 40
	fmt.Println(num)
	fmt.Printf("num type is %T\n", num)

	var defaultInt int
	fmt.Println(defaultInt)

	var str2 = "This is Go, but another string."
	fmt.Println(str2)
	fmt.Printf("str2 type is %T\n", str2)

	str3 := "This is Go, but again, another String"
	fmt.Println(str3)
	fmt.Printf("str3 type is %T\n", str3)
}
