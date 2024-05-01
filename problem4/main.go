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
func lrgstPlndrPrdct(fst, scnd int) int {
	max := 0
	for i := fst; i < 1000; i++ {
		for j := scnd; j < 1000; j++ {
			product := i * j
			if product > max && isPlndrm(product) {
				max = product
			}
		}

	}
	return max
}
func main() {
	max := lrgstPlndrPrdct(100, 100)
	fmt.Println(max)
}
