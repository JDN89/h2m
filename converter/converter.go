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
		// TODO: add curToken: int
		// TODO: add nextToken: int
	}
}

func (c *Converter) CollectTokens() {
	for tok := c.lexer.NextToken(); tok.Type != token.EOF; tok = c.lexer.NextToken() {
		c.tokens = append(c.tokens, tok)
	}
}

func (c *Converter) ConvertToMarkdown() []byte {

	converted := []byte{}

	for i := 0; i <= len(c.tokens)-1; i++ {

		tok := c.tokens[i]
		switch tok.Type {
		case token.CONTENT:
			contentSlice := c.lexer.Input[tok.StartPos:tok.EndPos]
			converted = append(converted, contentSlice...)
		case token.ANCHOR_OPEN:

			// TODO: peek next token. It should be content and if not throw an error for now
			converted = append(converted, '[')
			if href, ok := tok.Attributes["href"]; ok {
				converted = append(converted, []byte(href)...)
			}
			converted = append(converted, ']')
			// TODO: convert so that we go over tokens and also use peekToken. Now we just go over the tokens sequentially with a for loop, but I think it would be nicer if also here we keep calling advance and use a while statment for end not reach of converTokens bla bla bla. advance
			// I am used to working with C where we can use pointers. but here we don't have a pointer for the array so we still have to keep track of the indexes. so add them to the converter and increment them.

		default:

			converted = append(converted, []byte(GetMarkdown(tok.Type))...)
		}
		// ttypeString := c.tokens[i].Type.ToString()
		// fmt.Println(ttypeString)
	}

	return converted
}
