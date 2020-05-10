// Find maximum value in a slice

package main

import "fmt"

func main() {
	nums := []int{16, 23, 121, 14424, 2, 44}
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	fmt.Println(max)
}
