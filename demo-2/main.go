package main

import "fmt"

func main() {
	transactions := [5]int{12, 59, 31, 103, 95}
	transactionsPart := transactions[1:]
	transactionsPartNew := transactionsPart[:1]
	transactionsPartNew[0] = 30

	fmt.Println(transactions)
	fmt.Println(transactionsPartNew)
	fmt.Println(len(transactionsPart), cap(transactionsPart))
	fmt.Println(len(transactionsPartNew), cap(transactionsPartNew))
}
