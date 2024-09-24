package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	fmt.Println(states)

	states["CA"] = "California"
	states["OR"] = "Oregon"
	states["TX"] = "Texas"
	states["OH"] = "Ohio"
	states["FL"] = "Florida"
	states["AR"] = "Arizona"

	fmt.Println(states)

	// california := states["CA"]
	// fmt.Println(california)

	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println(keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
