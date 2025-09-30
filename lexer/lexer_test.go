package lexer

import (
	"h2m/token"
	"testing"
)

// func TestFixLexer(t *testing.T) {
// 	input := `<article> <h1> </h1> </article>`
//
// 	// expected token types and positions
// 	expected := []struct {
// 		typ      token.TokenType
// 		startPos int
// 		endPos   int
// 	}{
// 		{token.OPEN_ARTICLE, 0, 8},          // "<article>"
// 		{token.OPEN_HEADING_1, 10, 13},      // "<h1>"
// 		{token.CLOSED_HEADING_1, 15, 19},    // "</h1>"
// 		{token.CLOSED_ARTICLE, 21, 30},      // "</article>"
// 		{token.EOF, len(input), len(input)}, // end of input
// 	}
//
// 	l := New([]byte(input))
//
// 	for i, exp := range expected {
// 		tok := l.NextToken()
//
// 		if tok.Type != exp.typ {
// 			t.Errorf("token[%d] type mismatch: expected %s, got %s",
// 				i,
// 				token.TokenType.ToString(exp.typ),
// 				token.TokenType.ToString(tok.Type),
// 			)
// 		}
//
// 		if tok.StartPos != exp.startPos || tok.EndPos != exp.endPos {
// 			t.Errorf("token[%d] position mismatch: expected (%d,%d), got (%d,%d)",
// 				i, exp.startPos, exp.endPos, tok.StartPos, tok.EndPos,
// 			)
// 		}
// 	}
// }

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
