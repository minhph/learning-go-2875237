package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	// states holds a data struct consisiting of a hexadecimal
	// memory address of the object on heap, an integer (size), an integer(capacity)

	fmt.Println(states)
	states["WA"] = "Washington"
	states["NY"] = "New York"
	states["OL"] = "Oklahoma"
	states["MI"] = "Missisipi"
	fmt.Println(states)
	fmt.Printf("%p\n", states)

	oklahoma := states["OL"]
	fmt.Println(oklahoma)

	delete(states, "NY")
	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	// for...range is used for iteration, range is always used with for
	// this returns an index first, then a value one at a time

	keys := make([]string, len(states))
	fmt.Println(len(keys))

	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)
	sort.Strings(keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
