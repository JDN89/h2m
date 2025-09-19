package converter

import (
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

	l := lexer.New(input)
	c := New(l, []byte(input))
	c.CollectTokens()
	c.ConvertToMarkdown()
	fmt.Println("Done running TestConverter")
}
