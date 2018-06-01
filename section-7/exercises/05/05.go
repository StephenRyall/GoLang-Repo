package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := int64(100)

	a := big.NewInt(x)

	for x > 1 {
		b := big.NewInt(x - 1)
		a.Mul(a, b)
		x--
	}
	// a.Text(10)
	fmt.Println(a)
}
