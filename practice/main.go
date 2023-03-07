package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	colors = append(colors, "Black")
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1]) // return a subset of original arrays with index from a to b-1, exclusive b
	fmt.Println(colors)

	numbers := make([]int, 4) // capacity is not specified, which means it is default value, equal to the size
	numbers[0] = 12
	numbers[1] = 76
	numbers[2] = 92
	numbers[3] = 53

	numbers = append(numbers, 14, 315, 214, 1001, 234)
	// append can add more and more elements but Go have to reallocate if the size exceeds the capacity

	fmt.Println("Number string: ", numbers)
	fmt.Println(cap(numbers), len(numbers))
	fmt.Println("Before sort: ", &numbers[0], &numbers[1], &numbers[2])
	fmt.Printf("%p\n", &numbers)

	sort.Ints(numbers)
	fmt.Println("Number string: ", numbers)
	fmt.Println(cap(numbers), len(numbers))
	fmt.Println("After sort:", &numbers[0], &numbers[1], &numbers[2])
	// 0x1400001c0f0 0x1400001c0f8 0x1400001c100 continous 8-bits memory address

	fmt.Printf("%p\n", &numbers)
	// &number khác với &numbers[0] vì &numbers là địa chỉ của slice header trên stack, bao gồm capacity, size
	// còn numbers[0] là địa chỉ của element đầu tiên trong underlying array

}
