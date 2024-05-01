package main

import "fmt"

//нужно найти largest prime factor of 600851475143

func lrgstPrmFctr(num uint64) uint64 {
	var largest uint64
	// Find the largest prime factor
	// Continuously divide num by numbers starting from 2 up to num/2
	// Stop when num is no longer divisible or we reach 2
	// The last divisor is the largest prime factor
	for i := uint64(2); i*i <= num; i++ {
		for num%i == 0 {
			num /= i
			largest = i
		}
		if num > largest {
			largest = num
		}
	}
	return largest
}

func main() {
	fmt.Println(lrgstPrmFctr(600851475143))
}
