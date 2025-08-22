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
	input := `<article><header> content </header> </article>`

	expectedTokens := [5]token.TokenType{token.OPEN_ARTICLE, token.OPEN_HEADER, token.CLOSED_HEADER, token.CLOSED_ARTICLE, token.EOF}

	l := New(input)

	for i, expectedType := range expectedTokens {

		tok := l.NextToken()

		if tok.Type != expectedType {
			t.Errorf("tests compare token number[%d] - token type wrong. expected %s, got=%s", i, token.TokenType.ToString(expectedType), token.TokenType.ToString(tok.Type))
		}
	}
}

// func TestParseHeaderOfWhatColorIsYourFunction(t *testing.T) {
//
// 	input := ` <header>
//  <h1>
//    <a href="/2015/02/01/what-color-is-your-function/" rel="bookmark" title="Permanent Link to What Color is Your Function?">
//      What Color is Your Function?
//    </a>
//  </h1>
//    <a class="older" href="/2014/12/21/rooms-and-mazes/" title="Older Post “Rooms and Mazes: A Procedural Dungeon Generator”">←</a>
//    <a class="newer" href="/2015/09/07/what-the-hero-sees/" title="Newer Post “What the Hero Sees: Field-of-View for Roguelikes”">→</a>
//  </header>
// `
//
// }
