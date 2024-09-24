package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "hello from go!"
	file, err := os.Create("./fromString.txt")
	checkErr(err)
	length, err := io.WriteString(file, content)
	checkErr(err)
	fmt.Printf("Wrote a file with %v characters\n", length)
	defer file.Close()
	defer readFile("./fromString.txt")
}
func readFile(name string) {
	data, err := ioutil.ReadFile(name)
	checkErr(err)
	fmt.Println("Text from file: ", string(data))
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
