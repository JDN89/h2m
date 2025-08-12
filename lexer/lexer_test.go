package lexer

import (
	"h2m/token"
	"testing"
)

func TestGetNextToken(t *testing.T) {
	input := "<article>"

	expectedTokens := [2]token.TokenType{token.OPEN_ARTICLE, token.EOF}

	l := New(input)

	for i, expectedType := range expectedTokens {
		tok := l.NextToken()

		if tok.Type != expectedType {
			t.Errorf("tests[%d] - token type wrong. expected=%q, got=%q", i, expectedType, tok.Type)
		}
	}
}
