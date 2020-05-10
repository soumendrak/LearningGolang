// word counter
package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needless and pins
	Needless and pins
	Sew me a sail
	To catch me the wind
	`
	words := strings.Fields(text)
	fmt.Println(words)
	counts := map[string]int{}
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}
	fmt.Println(counts)
	// most common word
	max := 0
	mostCommon := ""
	for word, count := range counts {
		if count > max {
			mostCommon = word
			max = count
		}
	}
	fmt.Println("The most common word is", mostCommon)
}
