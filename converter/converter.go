package converter

import (
	"fmt"
	"h2m/lexer"
	"h2m/token"
)

type Converter struct {
	lexer  *lexer.Lexer
	tokens []token.Token
}

func New(l *lexer.Lexer) *Converter {
	return &Converter{
		lexer:  l,
		tokens: []token.Token{},
	}
}

func (c *Converter) CollectTokens() {
	for tok := c.lexer.NextToken(); tok.Type != token.EOF; tok = c.lexer.NextToken() {
		c.tokens = append(c.tokens, tok)
	}
}

func (c *Converter) advance() token.Token {
	return c.lexer.NextToken()
}

// TODO:
// Issie is that I want to receive the whole parsed file and then revert the token order and the convret to markdown
// Receive the whole input
// change it starting from the end
// instead of reverting the token array, i should just do length - i, so as not having to reverse the array which would be a costly action
func (c *Converter) ConvertToMarkdown() {

	i := len(c.tokens)
	for i = len(c.tokens) - 1; i >= 0; i-- {
		ttypeString := c.tokens[i].Type.ToString()
		fmt.Println(ttypeString)
	}
}
