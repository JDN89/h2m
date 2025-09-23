package lexer

import (
	"fmt"
	"h2m/token"
)

type Mode int

const (
	MODE_OUTSIDE_ARTICLE Mode = iota
	MODE_INSIDE_ARTICLE
	MODE_DONE
)

// TODO: make input private. Now public for debug purposes
type Lexer struct {
	Input    []byte
	startPos int
	currPos  int
	mode     Mode
}

func New(input []byte) *Lexer {
	l := &Lexer{Input: input, mode: MODE_OUTSIDE_ARTICLE}
	return l
}

func (l *Lexer) readChar() {
	l.currPos++
}

func (l *Lexer) currentChar() byte {
	return l.Input[l.currPos]
}

func (l *Lexer) consumeWhiteSpaceLineBreaks() {
	for {
		ch := l.currentChar()
		if ch != ' ' && ch != '\n' && ch != '\r' && ch != '\b' && ch != '\t' {
			break
		}
		l.readChar()
	}
}

func (l *Lexer) peekNextChar() byte {
	if l.currPos+1 == len(l.Input) {
		fmt.Printf("At end of input can't peek next character in the input stream! \n")
	}
	return l.Input[l.currPos+1]
}

// TODO : finish make anchor token
// func (l *Lexer) makeAnchortoken(token.TokenType) token.Token {
// 	for l.currentChar() != '>' {
// 		l.readChar()
// 	}
//
// 	//consume the whole anchor tag and convert to string
// 	slice := string(l.Input[l.startPos:l.currPos])
//
// 	if strings.Contains(slice, "href=") {
// 		strings.Split(slice, "href=")
// 	}
// 	strings.Split(slice)
//
// }

func (l *Lexer) makeHtmlElementToken(ttype token.TokenType) token.Token {

	// TODO: do I need a pointer? Try out at the end of project
	return token.Token{Type: ttype, StartPos: l.startPos, EndPos: l.currPos}
}

func (l *Lexer) makeContentToken() token.Token {
	for l.peekNextChar() != '<' {
		l.readChar()
	}
	end := l.currPos

	// consume '<'
	l.readChar()
	return token.Token{Type: token.CONTENT, StartPos: l.startPos, EndPos: end}
}

func (l *Lexer) consumeTag() token.Token {
	tok := token.Token{}

	switch l.currentChar() {

	case '/':
		l.readChar()
		switch l.currentChar() {
		case 'a':
			l.readChar()
			switch l.currentChar() {
			case 'r':
				tok = l.makeHtmlElementToken(token.CLOSED_ARTICLE)
			case '>':
				tok = l.makeHtmlElementToken(token.ANCHOR_CLOSED)
			}

		case 'h':
			l.readChar()
			switch l.currentChar() {
			case 'e':
				tok = l.makeHtmlElementToken(token.CLOSED_HEADER)
			case '1':
				tok = l.makeHtmlElementToken(token.CLOSED_HEADING_1)
			}
		}

	case 'a':

		l.readChar()
		switch l.currentChar() {
		case 'r':
			tok = l.makeHtmlElementToken(token.OPEN_ARTICLE)
		case ' ':
			tok = l.makeHtmlElementToken(token.ANCHOR_OPEN)

		}
	case 'h':
		l.readChar()
		switch l.currentChar() {
		case 'e':
			tok = l.makeHtmlElementToken(token.OPEN_HEADER)
		case '1':
			tok = l.makeHtmlElementToken(token.OPEN_HEADING_1)
		}
	}

	// consume '>'
	l.readChar()

	return tok
}

func (l *Lexer) NextToken() token.Token {
	tok := token.Token{}

	if l.currPos >= len(l.Input) {
		tok = token.Token{Type: token.EOF, StartPos: l.currPos}
		return tok
	}

	// up to the markdown converter to add tabs, enters,... where necessary
	l.consumeWhiteSpaceLineBreaks()

	l.startPos = l.currPos

	switch l.mode {

	case MODE_DONE:
		return token.Token{Type: token.EOF, StartPos: l.currPos}
	case MODE_INSIDE_ARTICLE:
		return l.lexInsideArticle()
	case MODE_OUTSIDE_ARTICLE:
		return l.lexOutsideArticle()
	default:
		panic(fmt.Sprintf("unexpected lexer.Mode: %#v", l.mode))
	}
}

func (l *Lexer) lexOutsideArticle() token.Token {
	for {
		switch l.currentChar() {
		case '<':
			l.startPos = l.currPos - 1
			l.readChar()
			tok := l.consumeTag()
			if tok.Type == token.OPEN_ARTICLE {
				l.mode = MODE_INSIDE_ARTICLE
				return tok
			}
			return tok
		case 0:
			return token.Token{Type: token.EOF, StartPos: l.currPos}
		default:
			l.readChar()
		}
	}
}

// Handles tokens between <article> ... </article>
func (l *Lexer) lexInsideArticle() token.Token {
	switch l.currentChar() {
	case '<':
		l.readChar()
		return l.consumeTag()
	case 0:
		return token.Token{Type: token.EOF, StartPos: l.currPos}
	default:
		return l.makeContentToken()
	}
}
