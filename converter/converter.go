package converter

import (
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

func (c *Converter) ConvertToMarkdown() []byte {

	converted := []byte{}

	for i := 0; i <= len(c.tokens)-1; i++ {

		tok := c.tokens[i]
		if tok.Type == token.CONTENT {
			sl := c.lexer.Input[tok.StartPos:tok.EndPos]
			converted = append(converted, sl...)
		} else {
			converted = append(converted, []byte(GetMarkdown(tok.Type))...)
		}
		// ttypeString := c.tokens[i].Type.ToString()
		// fmt.Println(ttypeString)
	}

	return converted
}
