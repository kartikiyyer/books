// My version of http request and getting header
package main

import (
	"fmt"
	"net/http"
)

func getHeader(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf("Unable to call GET")
	}
	defer resp.Body.Close()
	contentType := resp.Header.Get("Content-Type")

	if contentType == "" {
		return "", fmt.Errorf("content type is empty")
	}
	return contentType, nil
}

func main() {
	contentType, err := getHeader("https://golang.org")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(contentType)
	}
}
