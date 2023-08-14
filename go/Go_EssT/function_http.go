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
	
	contType := resp.Header.Get("Content-Type")
	if contType == "" {
		return "", fmt.Errorf("Content Type not found the response headers.")
	}

	return contType, nil
}

func main() {
	contType, err := contentType("https:/golang.org")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(contType)
	}
	contType, err = contentType("https://golang.org")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(contType)
	}

}