package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	/*
		type Calculator struct {
			operator1 string
			operator2 string
		}

		func (c Calculator) sum() {

		}
	*/

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Value 1: ")
	input1, _ := reader.ReadString('\n')
	operand1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}

	fmt.Print("Value 2: ")
	input2, _ := reader.ReadString('\n')
	operand2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 2 must be a number")
	}

	fmt.Print("Select an operation (+ - * /):")
	operator, _ := reader.ReadString('\n')
	operator = strings.TrimSpace(operator)
	switch operator {
	case "+":
		result := operand1 + operand2
		fmt.Printf("The result is: %.2f\n", result)
	case "-":
		result := operand1 - operand2
		fmt.Printf("The result is: %.2f\n", result)
	case "*":
		result := operand1 * operand2
		fmt.Printf("The result is: %.2f\n", result)
	case "/":
		result := operand1 / operand2
		fmt.Printf("The result is: %.2f\n", result)
	default:
		panic("Invalid operation")
	}
}
