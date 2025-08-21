package lexer

import (
	"h2m/token"
	"testing"
)

func TestMakeErrorToken(t *testing.T) {
	input := "<article blog content"
	expectedTokens := [1]token.TokenType{token.ERROR_NO_CLOSING_TAG}
	l := New(input)

	for i, expectedType := range expectedTokens {
		tok := l.NextToken()
		if tok.Type != expectedType {
			t.Errorf("tests[%d] - token type wrong. expected=%q, got=%q", i, expectedType, tok.Type)
		}

	}
}

func TestGetNextToken(t *testing.T) {
	input := "<article> content </article>"

	expectedTokens := [3]token.TokenType{token.OPEN_ARTICLE, token.CLOSED_ARTICLE, token.EOF}

	l := New(input)

	for i, expectedType := range expectedTokens {

		tok := l.NextToken()

		if tok.Type != expectedType {
			t.Errorf("tests compare token number[%d] - token type wrong. expected=%d, got=%d", i, expectedType, tok.Type)
		}
	}
}
