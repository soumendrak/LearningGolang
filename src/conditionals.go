// Check out various conditional operators

package main

import "fmt"

func ifElseCondition(x int) (message string) {
	if x > 5 {
		message = "x is greater than 5"
	} else if x == 5 {
		message = "x is equal to 5"
	} else {
		message = "x is less than 5"
	}
	return
}

func switchCase(x int) string {
	var message string
	switch x {
	case 5:
		message = "x is equal to 5"
	}
	return message
}

func main() {
	x := 10 - 5
	msg := ifElseCondition(x)
	fmt.Println("%s", msg)
	msg = switchCase(x)
	fmt.Println("%s", msg)
}
