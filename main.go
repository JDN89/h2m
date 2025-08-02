package main

import (
	"fmt"
	"h2m/lexer"
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

	l := lexer.New(string(body))

	fmt.Println("Fetched HTML:")
	fmt.Println("input lexer", l.Input)
}
