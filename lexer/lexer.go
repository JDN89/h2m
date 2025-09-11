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
	Input    string
	startPos int
	currPos  int
	char     byte
	mode     Mode
}

func New(input string) *Lexer {
	l := &Lexer{Input: input, mode: MODE_OUTSIDE_ARTICLE}
	l.readChar() // prime first char
	return l
}

// ReadChar will load the curernt char in l.char and advance currPos to +1
func (l *Lexer) readChar() {
	// TODO clenaup, leave now for debuggin purposes
	// fmt.Printf("inside readchar function %c \n", l.char)
	if l.currPos >= len(l.Input) {
		// end of file reached
		l.char = 0
	} else {
		l.char = l.Input[l.currPos]
	}
	l.currPos++
}

func (l *Lexer) consumeWhiteSpaceLineBreaks() {
	for {
		ch := l.char
		if ch != ' ' && ch != '\n' && ch != '\r' && ch != '\b' && ch != '\t' {
			break
		}
		l.readChar()
	}
}

func (l *Lexer) peekCurrPos() byte {
	if l.currPos == len(l.Input) {
		fmt.Printf("At end of input can't peek next character in the input stream! \n")
	}
	return l.Input[l.currPos]
}

func (l *Lexer) makeHtmlElementToken() token.Token {
	ref := l.Input[l.startPos:l.currPos]
	var ttype = token.HtmlReferenceTokenMap[string(ref)]
	// TODO do I need a pointer? Try out at the end of project
	return token.Token{Type: ttype, StartPos: l.startPos, EndPos: l.currPos - 1}
}

func (l *Lexer) makeContentToken() token.Token {
	start := l.currPos
	for l.peekCurrPos() != '<' {
		l.readChar()
	}
	end := l.currPos - 1
	// consume '<' so it gets loaded in lastConsumedChar
	l.readChar()
	return token.Token{Type: token.CONTENT, StartPos: start, EndPos: end}
}

func (l *Lexer) consumeTag() token.Token {
	tok := token.Token{}

	for l.char != '>' {
		l.readChar()
	}

	tok = l.makeHtmlElementToken()

	// consume '>' and load next char into lexer.char
	l.readChar()

	return tok
}

func (l *Lexer) NextToken() token.Token {

	tok := token.Token{}

	if l.currPos >= len(l.Input) {
		tok = token.Token{Type: token.EOF, StartPos: l.currPos}
		return tok
	}

	// See log 03/09/2025. whitespace, tabs,newlines, ... are not considered as content.
	// Otherwise for 2 tags seperated by a newline the lexer will consider the newline as content
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
		switch l.char {
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
	switch l.char {
	case '<':
		l.startPos = l.currPos - 1
		l.readChar()
		if l.char == '/' {
			l.readChar()
			tok := l.consumeTag()
			if tok.Type == token.CLOSED_ARTICLE {
				l.mode = MODE_DONE
			}
			return tok
		}
		return l.consumeTag()
	case 0:
		return token.Token{Type: token.EOF, StartPos: l.currPos}
	default:
		return l.makeContentToken()
	}
}
