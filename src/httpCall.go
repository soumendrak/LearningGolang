// make http call
package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	headers := resp.Header.Get("Content-Length")
	fmt.Println(resp.StatusCode, resp.Body)
	return headers, nil
}

func main() {
	url := "https://google.com"
	headers, err := contentType(url)
	if err != nil {
		fmt.Println("Error happened while fetching contentType")
	} else {
		fmt.Println("Content fetched successfully", headers)
	}

}
