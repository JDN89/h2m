package lexer

import (
	"h2m/token"
	"testing"
)

func TestGetNextToken(t *testing.T) {
	input := `<article> 		<header>
	<h1>
   <a href="/2015/02/01/what-color-is-your-function/" rel="bookmark" title="Permanent Link to What Color is Your Function?">
     What Color is Your Function?
   </a>
	</h1>
	</header> </article>`

	expectedAnchorTokenAttributes := map[string]string{
		"href":  "/2015/02/01/what-color-is-your-function/",
		"rel":   "bookmark",
		"title": "Permanent Link to What Color is Your Function?",
	}

	expectedTokens := [10]token.TokenType{token.OPEN_ARTICLE, token.OPEN_HEADER, token.OPEN_HEADING_1, token.ANCHOR_OPEN, token.CONTENT, token.ANCHOR_CLOSED, token.CLOSED_HEADING_1, token.CLOSED_HEADER, token.CLOSED_ARTICLE, token.EOF}

	l := New([]byte(input))

	for i, expectedType := range expectedTokens {

		tok := l.NextToken()

		if tok.Type != expectedType {
			t.Errorf("tests compare token number[%d] - token type wrong. expected %s, got=%s", i, token.TokenType.ToString(expectedType), token.TokenType.ToString(tok.Type))
		}
		if tok.Type == token.ANCHOR_OPEN {

			if len(tok.Attributes) != len(expectedAnchorTokenAttributes) {
				t.Fatalf("number of elements in attributes map doesn't match. We expected %d , but got %d", len(expectedAnchorTokenAttributes), len(tok.Attributes))
			}

			for attr, content := range expectedAnchorTokenAttributes {
				if tok.Attributes[attr] != content {
					t.Fatalf("Expected %s, but got %s ", content, tok.Attributes[attr])
				}
			}

		}

	}
}
