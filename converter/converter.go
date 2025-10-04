package converter

import (
	"fmt"
	"h2m/lexer"
	"h2m/token"
)

type Converter struct {
	lexer *lexer.Lexer
	//TODO: convert to go chan . so the converter owns a go channel
	tokens    []token.Token
	pos       int // position of nextToken
	currToken token.Token
	nextToken token.Token
}

func New(l *lexer.Lexer) *Converter {
	c := Converter{
		lexer:  l,
		tokens: []token.Token{},
		pos:    0,
	}

	c.CollectTokens()
	c.currToken = c.tokens[c.pos]
	c.pos++
	c.nextToken = c.tokens[c.pos]
	return &c
}

func (c *Converter) CollectTokens() {

	// infinit for loop. similar to while(;;) in C. Just keep looping until you enter the EOF token
	for {
		tok := c.lexer.NextToken()

		c.tokens = append(c.tokens, tok)
		if tok.Type == token.EOF {
			break
		}

	}
}

func (c *Converter) advance() {
	c.pos++
	c.currToken = c.nextToken
	c.nextToken = c.tokens[c.pos]
}

func (c *Converter) ConvertToMarkdown() []byte {

	markdown := []byte{}

	for c.nextToken.Type != token.EOF {

		tok := c.currToken

		switch tok.Type {
		case token.CONTENT:
			contentSlice := c.lexer.Input[tok.StartPos:tok.EndPos]
			markdown = append(markdown, contentSlice...)
		case token.ANCHOR_OPEN:

			markdown = append(markdown, '[')
			if href, ok := tok.Attributes["href"]; ok {
				markdown = append(markdown, []byte(href)...)
			} else {
				fmt.Printf("The ANCHOR_OPEN token didn't contain a href key value pair! \n")
			}
			markdown = append(markdown, ']')
			if c.nextToken.Type == token.CONTENT {
				// consume the currentToken
				c.advance()
				markdown = append(markdown, '(')
				// the endPos in contentToken is right before the '<'. in slice [x:y], the y value is not included. So + 1
				contentSlice := c.lexer.Input[c.currToken.StartPos : c.currToken.EndPos+1]
				markdown = append(markdown, contentSlice...)
				markdown = append(markdown, ')')
			} else {
				fmt.Printf("Expected Content token but got: %s \n", token.TokenType.ToString(c.nextToken.Type))
			}

		default:
			markdown = append(markdown, []byte(GetMarkdown(tok.Type))...)
		}
		c.advance()
	}
	return markdown
}
