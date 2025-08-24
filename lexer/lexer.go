package lexer

import (
	"fmt"
	"h2m/token"
)

// TODO make input private. Now public for debug purposes
type Lexer struct {
	Input    string
	startPos int
	currPos  int
	char     byte
	//TODO think of better way to start and stop lexing
	start bool
	stop  bool
}

func New(input string) *Lexer {
	l := &Lexer{Input: input}
	// load first char into the lexer. I was first doing this at the beginning of NextToken, but this breaks when we start parsing the second token, because < is allready loaded into the lexer
	l.readChar()
	return l
}

// ReadChar will load the curernt char in l.char and advance currPos to +1
// TODO probably have to rename currPos
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

func (l *Lexer) peekCurrent() byte {
	return l.char
}

func (l *Lexer) consumeWhiteSpaceLineBreaks() {
	for {
		ch := l.peekCurrent()
		if ch != ' ' && ch != '\n' && ch != '\r' {
			break
		}
		l.readChar()
	}
}

func (l *Lexer) peekNext() byte {
	if l.currPos == len(l.Input) {
		fmt.Printf("At end of input can't peek next character in the input stream! \n")
	}
	return l.Input[l.currPos]
}

func isChar(char byte) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z'
}

// NOTE: Token should contain a proper errorMessage. But I want to keep my tokens as lightweight as possilbe.
// Use different Error tokens, and based on the type of error token and position of  the error I can give a clear error message
func (l *Lexer) makeErrorToken() token.Token {
	return token.Token{Type: token.ERROR_NO_CLOSING_TAG, Pos: l.currPos}
}

func (l *Lexer) makeHtmlElementToken() token.Token {
	ref := l.Input[l.startPos:l.currPos]
	var ttype = token.HtmlReferenceTokenMap[string(ref)]
	// TODO do I need a pointer? Try out at the end of project
	return token.Token{Type: ttype, Pos: l.startPos}
}

// NOTE Only make tokens for the html tags, no content token needed, because no manipulation of content.
func (l *Lexer) consumeTag() token.Token {
	tok := token.Token{}

	l.consumeWhiteSpaceLineBreaks()

	for isChar(l.char) {
		l.readChar()
	}

	if l.peekCurrent() != '>' {
		return l.makeErrorToken()
	}

	tok = l.makeHtmlElementToken()
	//consume '>'
	l.readChar()
	return tok
}

// BUG: I have create tokens for the content as well, otherwise when i parse <article> bla bla </article>
// lexer will return a token for <article> for now just consumes the content and sets l.start
func (l *Lexer) NextToken() token.Token {

	tok := token.Token{}

	if l.currPos >= len(l.Input) {
		tok = token.Token{Type: token.EOF, Pos: l.currPos}
		return tok
	}

	l.startPos = l.currPos

	// as long start == false keep reading char
	// check encounter <article>
	// set start == true

	// NOTE no pointer because the struct is very small.
	// TODO test impact of using pointer once project is finished

	// Parsed </article>
	for l.start == false && l.stop == true {
		l.readChar()
	}

	for l.start == false && l.stop == false {
		switch l.char {

		// NOTE consume until you reach <article>
		//TODO what to do after we parsed </article>
		case '<':
			// NOTE: because I don't make a token for content, we consume the content with the token that follows, which means we have to set the l.start when we arrive at '<'
			// -1 because once char is and set, currPos allready points to the next char to be consumed
			l.startPos = l.currPos - 1
			l.readChar()
			//NOTE we consume the <tag> and make a token
			tok = l.consumeTag()
			// start lexing content of the blog article
			if tok.Type == token.OPEN_ARTICLE {
				l.start = true
				return tok
			}
			// in case we don't find a closing tag for the html tag
			if tok.Type == token.ERROR_NO_CLOSING_TAG {
				return tok
			}

		// Stop parsing when we reach EOF, in that case we set char in readchar to 0 and hit this case
		case 0:
			tok = token.Token{Type: token.EOF, Pos: l.currPos}
			return tok

		default:
			l.readChar()
		}

	}
	for l.start == true && l.stop == false {
		switch l.char {
		case '<':
			// -1 because l.currPos allreay points to the next char in the input
			l.startPos = l.currPos - 1
			l.readChar()
			if l.peekCurrent() == '/' {
				//BUG set breakpoint here
				l.readChar()
				tok := l.consumeTag()
				if tok.Type == token.CLOSED_ARTICLE {
					l.stop = true
					l.start = false
				}
				return tok
			}
			tok := l.consumeTag()
			return tok

			//NOTE  EOF we set byte to 0 in readchar 0x00 when we reach the end of the input
			// I could also stop lexing at </article>, but I'll just keep reading the leftover tags
		case 0:
			tok = token.Token{Type: token.EOF, Pos: l.currPos}
			return tok
		default:
			// TODO for now I consume al chars. Should I make a difference between char, whitespace and other symbols?
			l.readChar()
		}
	}
	return tok
}
