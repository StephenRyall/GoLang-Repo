package main

import "fmt"

func main() {
	transactions := make([][]int, 0)

	for i := 4; i < 16; i++ {
		transaction := make([]int, 0)
		for j := 0; j < 8; j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println(transactions)
}
