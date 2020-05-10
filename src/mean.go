// Calculate the mean of two provided numbers
package main

import (
	"fmt"
	"os"
	"strconv"
)

func mean(a float64, b float64) float64 {
	return (a + b) / 2
}

func main() {
	x, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Errorf("Valid command line arguments not provided")
	}
	y, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Errorf("Valid command line arguments not provided")
	}
	mean := mean(x, y)
	fmt.Printf("mean=%v\n", mean)
}
