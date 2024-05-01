package main

import "fmt"

func multpls(maxNum uint64) uint64 {

	var sum uint64
	var i uint64

	for i = 1; i < maxNum; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	fmt.Println(multpls(1000))
}
