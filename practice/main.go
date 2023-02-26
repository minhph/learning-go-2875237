package main

import (
	"fmt"
)

func main() {
	// This is a struct
	type Hero struct {
		FirstSkill  string
		SecondSkill string
	}

	rouie := Hero{"Holy Emblem", "Holy Gate"}
	fmt.Println(rouie)
}
