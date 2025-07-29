package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	res, err := http.Get("https://journal.stuffwithstuff.com/2015/02/01/what-color-is-your-function/")
	if err != nil {
		fmt.Println("Error fetching url:", err)
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading contents of response body")
	}
	fmt.Println("Fetched HTML:")
	fmt.Println(string(body))
}
