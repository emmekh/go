package main

import "fmt"

func main() {
	num := make(chan int)
	numSq := make(chan int)

	go func() {
		for x := 0; x <= 10; x++ {
			num <- x
		}
		close(num)
	}()

	go func() {
		for x := range num {
			numSq <- x * x
		}
		close(numSq)
	}()

	for x := range numSq {
		fmt.Println(x)
	}

}
