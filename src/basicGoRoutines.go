// Basic concurrent goroutines execution.
// Command to run:
// time go run basicGoRoutines.go
package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.fb.com",
	}
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			returnType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

// this is with sync group
func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Println(url, "->", ctype)
}
