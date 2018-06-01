package main

import "fmt"

func main () {
	var input int32
	fmt.Printf("Enter integer: ")
	fmt.Scan(&input)

	if input % 2 == 0 {
		newNum := input/2
		fmt.Printf("Half of input: %g\nStatus: %v\n", newNum, true)
	} else {
		newNum := input/2
		fmt.Printf("Half of input: %.g\nStatus: %v\n", newNum, false)
	}
}

