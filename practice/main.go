package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	fmt.Println("Time now is: ", n)

	t := time.Date(2023, time.April, 12, 18, 30, 30, 30, time.UTC)
	timeFormatted := t.Format(time.ANSIC)
	fmt.Println("I was born at", timeFormatted)

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00:00 2009")
	fmt.Printf("The type of parsed Times is: %T", parsedTime)
}
