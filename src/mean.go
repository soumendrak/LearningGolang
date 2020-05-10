// Calculate the mean of two provided numbers
package main

import (
	"fmt"
	"os"
)

func mean(a float64, b float64) float64 {
	return (a + b) / 2
}

func main() {
	x, y := 1.0, 2.0 // TODO: read these from command line arguments
	fmt.Println(os.Args)
	mean := mean(x, y)
	fmt.Printf("mean=%v\n", mean)
}
