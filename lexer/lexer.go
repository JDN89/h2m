package lexer

import (
	"fmt"
	"h2m/token"
	"regexp"
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

func (l *Lexer) NextToken() token.Token {
	// if for some reason </article> was the last html tag. Then we don't immediatly hit MODE DONE
	// MODE done is to exit early when we parsed the closing article html tag
	// In my test cases </article> is often the last token, so we hit this if condition before we reach MODE_DONE
	// My test cases will be more realistic once i parse the whole blog
	if l.currPos >= len(l.Input) {
		return token.Token{Type: token.EOF, StartPos: l.currPos, EndPos: l.currPos}
	}

	// up to the markdown converter to add tabs, enters,... where necessary
	l.consumeWhiteSpaceLineBreaks()

	l.startPos = l.currPos

	switch l.mode {

	case MODE_DONE:
		return token.Token{Type: token.EOF, StartPos: l.startPos, EndPos: l.currPos}
	case MODE_INSIDE_ARTICLE:
		return l.lexInsideArticle()
	case MODE_OUTSIDE_ARTICLE:
		return l.lexOutsideArticle()
	default:
		panic(fmt.Sprintf("unexpected lexer.Mode: %#v", l.mode))
	}
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

func (l *Lexer) makeAnchortoken(ttype token.TokenType) token.Token {
	for l.currentChar() != '>' {
		l.readChar()
	}

	//consume the whole anchor tag and convert to string
	anchorSlice := string(l.Input[l.startPos:l.currPos])

	hrefRegexPatterns := map[string]*regexp.Regexp{

		"href":  regexp.MustCompile(`href="([^"]+)"`),
		"rel":   regexp.MustCompile(`rel="([^"]+)"`),
		"title": regexp.MustCompile(`title="([^"]+)"`),
	}

	hrefAttributes := make(map[string]string)

	// FindSubmatch returns the whole match and then the capture groups if any are found
	for key, regex := range hrefRegexPatterns {
		attributeMatch := regex.FindSubmatch([]byte(anchorSlice))
		if len(attributeMatch) > 1 {
			hrefAttributes[key] = string(attributeMatch[1])
		}

	}

	return token.Token{Type: ttype, StartPos: l.startPos, EndPos: l.currPos, Attributes: hrefAttributes}

}

func (l *Lexer) makeHtmlElementToken(ttype token.TokenType) token.Token {

	for l.currentChar() != '>' {
		l.readChar()
	}
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
			tok = l.makeAnchortoken(token.ANCHOR_OPEN)

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

func (l *Lexer) lexOutsideArticle() token.Token {
	for {
		switch l.currentChar() {
		case '<':
			l.readChar()
			tok := l.consumeTag()
			if tok.Type == token.OPEN_ARTICLE {
				l.mode = MODE_INSIDE_ARTICLE
				return tok
			}
			return tok
		// case 0:
		// 	return token.Token{Type: token.EOF, StartPos: l.currPos}
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
		tok := l.consumeTag()
		if tok.Type == token.CLOSED_ARTICLE {
			l.mode = MODE_DONE
		}
		return tok

	// case 0:
	// 	return token.Token{Type: token.EOF, StartPos: l.currPos}
	default:
		return l.makeContentToken()
	}
}
