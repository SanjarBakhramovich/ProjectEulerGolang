// Smallest Multiple

package main

import "fmt"

func smllstMultple(num int) int {
	for i := num; ; i++ {
		if i%11 == 0 && i%12 == 0 && i%13 == 0 && i%14 == 0 && i%15 == 0 && i%16 == 0 && i%17 == 0 && i%18 == 0 && i%19 == 0 && i%20 == 0 {
			return i
		}
	}
}

func main() {
	fmt.Println(smllstMultple(20))
}
