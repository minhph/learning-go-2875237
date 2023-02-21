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
	input, err1 := reader.ReadString('\n')
	fmt.Println("You entered:", input)

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("no error")
	}

	fmt.Print("Enter a number: ")
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number:", aFloat)
	}

	// reader := bufio.NewReader(os.Stdin): This creates a new buffered reader that reads
	// from the standard input (i.e. the console).
	// The bufio package is used for buffered input/output operations,
	// which can be more efficient than reading or writing one byte at a time.

}
