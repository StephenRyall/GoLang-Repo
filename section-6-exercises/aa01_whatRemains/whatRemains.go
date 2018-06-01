package main

import "fmt"

func main () {
	var largeNum int32
	var smallNum int32

	fmt.Print("Please enter a small number: ")
	fmt.Scan(&smallNum)
	fmt.Print("Please enter a large number: ")
	fmt.Scan(&largeNum)

	remainder := largeNum % smallNum

	fmt.Printf("The remainder is: %v \n", remainder)
}