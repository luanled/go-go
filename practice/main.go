package main

import (
	"fmt"
)

func main() {
	// poodle := Dog{"Poodle", 10, "Woof!"}
	// fmt.Println(poodle)
	// fmt.Printf("%+v\n", poodle)
	// fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
	// poodle.Speak()
	// poodle.SpeakLikeCrazy()
	fmt.Println(addValues(1.55, 3.65))
}

// Dog is a struct
type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

func (d Dog) SpeakLikeCrazy() {
	d.Sound = fmt.Sprintf("%v %v %v %v %v %v %v %v %v %v", d.Sound, d.Sound, d.Sound, d.Sound, d.Sound, d.Sound, d.Sound, d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}
