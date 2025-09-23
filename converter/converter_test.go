package converter

import (
	"bytes"
	"fmt"
	"h2m/lexer"
	"testing"
)

func TestConverter(t *testing.T) {

	input := `<article> 		<header>
	<h1>
	  <a href="/2015/02/01/what-color-is-your-function/" rel="bookmark" title="Permanent Link to What Color is Your Function?">
	    What Color is Your Function?
	  </a>
	</h1>
	</header> </article>`

	output := `
# [What Color is Your Function?](/2015/02/01/what-color-is-your-function/ "Permanent Link to What Color is Your Function?")`

	l := lexer.New([]byte(input))
	c := New(l)
	c.CollectTokens()
	result := c.ConvertToMarkdown()
	if !bytes.Equal(result, []byte(output)) {
		fmt.Printf("result -- %s", result)
	}
	c.ConvertToMarkdown()

	fmt.Printf("result -- %s", result)
	fmt.Println("Done running TestConverter")
}
