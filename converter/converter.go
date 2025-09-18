package converter

import (
	"fmt"
	"h2m/lexer"
	"h2m/token"
)

type Converter struct {
	lexer  *lexer.Lexer
	tokens []token.Token
	input  []byte
}

// TODO : fix
func New(l *lexer.Lexer) *Converter {
	return &Converter{
		currentToken: l.NextToken(),
		lexer:        l,
	}
}

// TODO: Collect the tokens. while not EOF append to tokens array

func (c *Converter) advance() {
	c.currentToken = c.lexer.NextToken()
}

// TODO:
// Issie is that I want to receive the whole parsed file and then revert the token order and the convret to markdown
// Receive the whole input
// change it starting from the end
// instead of reverting the token array, i should just do length - i, so as not having to reverse the array which would be a costly action
func (c *Converter) ConvertToMarkdown() {
	for c.currentToken.Type != token.EOF {
		fmt.Printf("token %s\n", c.currentToken.Type.ToString())
		c.advance()
	}
}
