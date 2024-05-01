package main

import "fmt"

//нужно найти сумму чётных чисел ряда Фибоначчи до 4000000

func fibSumEven(limit uint32) uint32 {
	var fib1, fib2, sumEvens uint32 = 1, 2, 0

	for fib2 < limit {
		if fib2%2 == 0 {
			sumEvens += fib2
		}
		fib1, fib2 = fib2, fib1+fib2
	}
	return sumEvens
}

func main() {
	fmt.Println(fibSumEven(4000000))
}
