package main

import (
	"fmt"
	"math"
)

func main() {
	i1, i2, i3 := 22, 45, 59
	intSum := i1 + i2 + i3
	fmt.Println(intSum)

	f1, f2, f3 := 22.34, 421.24, 120.31
	floatSum := f1 + f2 + f3
	fmt.Println(floatSum)

	floatSum = math.Round(floatSum)
	fmt.Println("The sum is now: ", floatSum)
}
