package main

import "fmt"

func biggest(number ...int) (res int) {

	for i := 0; i < len(number); i++ {
		if number[i] > res {
			res = number[i]
		}
	}
	return res
}

func main() {
	fmt.Println(biggest(1, 2, 56, 24, 3, 66, 7))
}
