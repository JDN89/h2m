package main

import (
	"fmt"
	"h2m/converter"
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

	l := lexer.New(body)
	c := converter.New(l)
	// TODO: I want to optimize it later. CollectTokens concrentyl and per thread then ConvertToMarkdown, and then stitch the whole back together based on there positions
	c.CollectTokens()
	c.ConvertToMarkdown()
	fmt.Println("DONE")
}
