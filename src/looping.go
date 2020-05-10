// Check for loops
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("The current value of i is: ", i)
		}
	}
	fmt.Println("----")
	b := 0
	for {
		if b > 6 {
			break
		}
		fmt.Println(b)
		b++
	}
}
