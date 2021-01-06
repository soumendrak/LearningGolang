package main

import "fmt"

func main() {
	NewMap := map[string]int{
		"a": 123,
		"b": 234,
	}
	if val, ok := NewMap["f"]; ok {
		fmt.Println("The element is found: ", val)
	} else {
		fmt.Println("The value is not present")
	}
}
