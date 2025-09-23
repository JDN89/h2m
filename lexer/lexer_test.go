package lexer

import (
	"h2m/token"
	"testing"
)

// TODO: finish make anchor token
func testLexAnchorToken(t *testing.T) {

	input := `    <a href="/2015/02/01/what-color-is-your-function/" rel="bookmark" title="Permanent Link to What Color is Your Function?">
   `

	expectedAnchorToken := token.Token{
		Type:     token.ANCHOR_OPEN,
		StartPos: 1,
		EndPos:   108,
		Attributes: map[string]string{
			"href":  "/2015/02/01/what-color-is-your-function/",
			"rel":   "bookmark",
			"title": "Permanent Link to What Color is Your Function?",
		},
	}

	l := New([]byte(input))
	result := l.NextToken()

	if result.Type != expectedAnchorToken.Type {
		t.Errorf("tests compare token number[%d] - token type wrong. expected %s, got=%s", token.TokenType.ToString(expectedAnchorToken.Type), token.TokenType.ToString(result.Type))
	}

}

func TestGetNextToken(t *testing.T) {
	input := `<article> 		<header>
	<h1>
   <a href="/2015/02/01/what-color-is-your-function/" rel="bookmark" title="Permanent Link to What Color is Your Function?">
     What Color is Your Function?
   </a>
	</h1>
	</header> </article>`

	expectedTokens := [10]token.TokenType{token.OPEN_ARTICLE, token.OPEN_HEADER, token.OPEN_HEADING_1, token.ANCHOR_OPEN, token.CONTENT, token.ANCHOR_CLOSED, token.CLOSED_HEADING_1, token.CLOSED_HEADER, token.CLOSED_ARTICLE, token.EOF}

	l := New([]byte(input))

	for i, expectedType := range expectedTokens {

		tok := l.NextToken()

		if tok.Type != expectedType {
			t.Errorf("tests compare token number[%d] - token type wrong. expected %s, got=%s", i, token.TokenType.ToString(expectedType), token.TokenType.ToString(tok.Type))
		}
	}
}
