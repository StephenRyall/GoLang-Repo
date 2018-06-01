package main

import (
	"fmt"
)

func PrimeFactors(n int) int {
	var pfs []int
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}
	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}
	if n > 2 {
		pfs = append(pfs, n)
	}

	return pfs[len(pfs)-1]
}

func main() {
	for i := 2; i < 200; i++ {
		fmt.Println(PrimeFactors(i))
	}

}
