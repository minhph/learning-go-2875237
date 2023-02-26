package main

import (
	"fmt"
)

func main() {
	printSomething()
	sumTwoValues := addTwoValues(3, 4)
	sumManyValues := addManyValues(84, 512, 53)
	fmt.Println(sumTwoValues)
	fmt.Println(sumManyValues)
}

func printSomething() {
	fmt.Println("Print something.")
}

func addTwoValues(value1, value2 int) int {
	return value1 + value2
}

func addManyValues(values ...int) int {
	total := 0
	for v := range values {
		total = total + values[v]
	}
	return total
}
