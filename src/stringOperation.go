package main

import (
	"fmt"
	"strings"
)

func main() {
	var str strings.Builder
	for i := 0; i < 10; i++ {
		str.WriteString("a")
	}
	fmt.Println(str.String())
}

func init() {
	fmt.Println("This function is called first")
}
