// Check even ended numbers
// i.e. number with same first and last digit

package main

import (
	"fmt"
)

func main() {
	checkEvenEndedNumbers()
}

func checkEvenEndedNumbers() {
	count := 0
	for i := 1; i < 9; i++ {
		for j := i; j < 9; j++ {
			number := fmt.Sprintf("%d", i*j)
			if number[0] == number[len(number)-1] {
				fmt.Println("Even Ended Numbers")
				count++
				fmt.Println(number[0], number[len(number)-1], number, count)
			} else {
				fmt.Println("NOT Even Ended Numbers")
			}
		}
	}
}
