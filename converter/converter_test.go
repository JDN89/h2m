package converter

import (
	"h2m/lexer"
	"strings"
	"testing"
)

func normalize(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func TestConverter(t *testing.T) {

	input := `<article> 		<header>
	<h1>
	  <a href="/2015/02/01/what-color-is-your-function/" rel="bookmark" title="Permanent Link to What Color is Your Function?">
	    What Color is Your Function?</a>
	</h1>
	</header> </article>`

	expected := `# [/2015/02/01/what-color-is-your-function/](What Color is Your Function?)`

	l := lexer.New([]byte(input))
	c := New(l)
	c.CollectTokens()
	result := string(c.ConvertToMarkdown())

	if normalize(result) != normalize(expected) {
		t.Fatalf("result : \n%s\nexpected :\n%s\n", result, expected)
	}

	//NOTE: comparing bytes was very error prone. as soon as there was a space or new line to many, the bytes didn't match. Solved this by converting to strings and trimming the redundant " ", '/n', '/t' ,... https://stackoverflow.com/questions/37290693/how-to-remove-redundant-spaces-whitespace-from-a-string-in-golang?utm_source=chatgpt.com

	// if !bytes.Equal(result, []byte(output)) {
	// 	t.Fatalf("result : \n %s \n output :  \n %s \n", result, output)
	// }

	t.Log("Done running TestConverter")
}
