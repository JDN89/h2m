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
	start    bool
}

func New(input string) *Lexer {
	l := &Lexer{Input: input}
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

func (l *Lexer) peekNext() byte {
	if l.currPos == len(l.Input) {
		fmt.Printf("At end of input can't peek next character in the input stream! \n" )
	}
	return l.Input[l.currPos]
}

func isChar(char byte) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z'
}

// NOTE: Token should contain a proper errorMessage. But I want to keep my tokens as lightweight as possilbe.
// Use different Error tokens, and based on the type of error token and position of  the error I can give a clear error message
func (l *Lexer) makeErrorToken(errorMessage string) token.Token {
	return token.Token{Type: token.ERROR_NO_CLOSING_TAG, Pos: l.currPos}
}

func (l *Lexer) makeHtmlElementToken() token.Token {
	ref := l.Input[l.startPos:l.currPos]
	var ttype = token.HtmlReferenceTokenMap[string(ref)]
	// TODO do I need a pointer? Try out at the end of project
	return token.Token{Type: ttype, Pos: l.startPos}
}

// NOTE Only make tokens for the html tags, no content token needed, because no manipulation of content.
func (l *Lexer) readIdentifier() token.Token {
	tok := token.Token{}
	for isChar(l.char) {
		l.readChar()
	}

	if l.peekCurrent() != '>' {
		return l.makeErrorToken("Expected '>' after identifier")
	}

	tok = l.makeHtmlElementToken()
	//consume '>'
	l.readChar()
	return tok
}

func (l *Lexer) NextToken() token.Token {

	l.startPos = l.currPos
	// Consume first token and load in to l.char
	l.readChar()

	// as long start == false keep reading char
	// check encounter <article>
	// set start == true

	// NOTE no pointer because the struct is very small.
	// TODO test impact of using pointer once project is finished
	tok := token.Token{}

	for l.start == false {
		switch l.char {

		// NOTE consume until you reach <article>
		//TODO what to do after we parsed </article>
		case '<':
			l.readChar()
			//NOTE we consume the <Identifier> and make a token
			tok = l.readIdentifier()
			// start lexing content of the blog article
			if tok.Type == token.OPEN_ARTICLE {
				l.start = true
				return tok
			}

		// TODO stop parsing if we reach end of file and never encounter article
		default:
			l.readChar()
		}

	}
	for l.start == true {
		// TODO: implement make tokens for the other element references;
		switch l.char {
		case '<':
			fmt.Println("< found")
			l.readChar()
		case '/':
			fmt.Println("/ found")
			l.readChar()
		case '>':
			fmt.Println("> found. Make token")
			l.readChar()

			//NOTE  EOF we set byte to 0 in readchar 0x00 when we reach the end of the input
			// I could also stop lexing at </article>, but I'll just keep reading the leftover tags
		case 0:
			tok = token.Token{Type: token.EOF, Pos: l.currPos}
			return tok
		default:
			if isChar(l.char) {
				fmt.Printf("char found %c \n", l.char)
				//NOTE I prefer to not call readchar in the isChar funciton,
				// so it's more explicit where we advance and return the next char
				l.readChar()
				break
			}
			fmt.Println("no char found")

		}

	}
	return tok
}

// latest TODO: <div>  make token when you encounter > or </
// debug and see why test doesn't print a char
//look at video of tj to see how to debug go code in neovim

//TODO define tokens
// Type and position in input stream
// no need to copy the string, calculate the offset
// identifier tokens <p> </p>
//content Token -> part between <p> </p>
//just track start position  end of contennt is start of enclosing tag

// NOTE: I noticed that the to blogs that I want to convert to markdown the blog content is wrapped in the <article> elmenent. So I'll start lexing once I encounter article. When encounter '<' identify which element it is and store the type and position in a token struct. for the content als just the position is needed. altough why? we just need to replace the html tags and not the content. So I can probably start with just replacing the tags.
//WHAT with tags that immediatley have a closing tag, tables, hrefs, lists? worries for later
