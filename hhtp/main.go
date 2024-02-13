package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://freshman.tech/")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("HTTP Status Code:", resp.StatusCode)
}
