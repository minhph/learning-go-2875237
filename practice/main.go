package main

import (
	"fmt"
)

func main() {
	anInt := 2
	var p *int // a pointer pointing to an int value, zero-value: nil
	p = &anInt // assign address of anInt to p (type stricted)
	fmt.Println(*p)
}
