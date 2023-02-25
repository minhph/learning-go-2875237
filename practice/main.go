package main

import (
	"fmt"
)

func main() {
	anInt := 2
	var p *int      // a pointer pointing to an int value, zero-value: nil
	p = &anInt      // assign address of anInt to p (type stricted)
	fmt.Println(p)  // the memory address which pointer is pointing to
	fmt.Println(*p) // the value stored at the memory address the pointer is pointing to
	fmt.Println(&p) // the memory address of the pointer itself
}
