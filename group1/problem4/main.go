package main

import (
	"fmt"
	"strconv"
)

//Largest Palindrome product
//Find the largest palindrome made from the product of two 3-digit numbers

func isPlndrm(num int) bool {
	str := strconv.Itoa(num)
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}
func lrgstPlndrPrdct() int {
	max := 0
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			product := i * j
			if product > max && isPlndrm(product) {
				max = product
			}
		}

	}
	return max
}
func main() {
	max := lrgstPlndrPrdct()
	fmt.Println(max)
}
