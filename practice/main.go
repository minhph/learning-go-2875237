package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Value 1: ")
	stringInput1, _ := reader.ReadString('\n')
	floatOperator1, err1 := strconv.ParseFloat(strings.TrimSpace(stringInput1), 64)

	if err1 != nil {
		fmt.Println(err1)
		panic("The input value must be a number.")
	}

	fmt.Print("Value 2: ")
	stringInput2, _ := reader.ReadString('\n')
	floatOperator2, err2 := strconv.ParseFloat(strings.TrimSpace(stringInput2), 64)

	if err2 != nil {
		fmt.Println(err2)
		panic("The input value must be a number.")
	}

	sum := floatOperator1 + floatOperator2
	fmt.Printf("The sum of %.2f and %.2f is %.2f\n", floatOperator1, floatOperator2, sum)

	// The ReadString method actually invokes standard input (console).
	// The NewReader(os.Stdin) method only return an reader object that can read from standard input.
	// It does not invoke input.
	// Therefore, we only need 1 reader for this.
	// The argument "\n" on reader.ReadString('\n') means saves characters on buffer memory until the delimiter "\n" is reached.
}
