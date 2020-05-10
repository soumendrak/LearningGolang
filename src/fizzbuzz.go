// fizzbuzz problem solution

package main

import "fmt"

func main() {
	min, max := 0, 20
	fizzbuzz(min, max)
}

func fizzbuzz(min, max int) {
	for i := min; i < max; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
