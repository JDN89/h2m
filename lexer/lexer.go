package lexer

import (
	"fmt"
	"h2m/token"
	"log"
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
	//load one char in l.char
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	fmt.Printf("inside readchar function %c \n", l.char)
	if l.currPos >= len(l.Input) {
		l.char = 0
	} else {
		l.char = l.Input[l.currPos]
	}
	l.currPos++
}

func (l *Lexer) peek() byte {
	//BUG here
	// if l.currPos == len(l.Input) {
	// 	fmt.Printf("At end of input %c can't peek! \n", l.Input[l.currPos])
	// }
	return l.char
}

func (l *Lexer) peekNext() byte {
	if l.currPos == len(l.Input) {
		fmt.Printf("At end of input %c can't peek! \n", l.Input[l.currPos])
	}
	return l.Input[l.currPos]
}

func isChar(char byte) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z'
}

func (l *Lexer) makeToken() token.Token {
	ref := l.Input[l.startPos:l.currPos]
	var o string = string(ref)
	var ttype = token.HtmlReferenceTokenMap[o]
	// TODO do I need a pointer?
	return token.Token{Type: ttype, Pos: l.startPos}
}

// NOTE I am not making tokens for the content between the tags, because I can assume the start and end based on the html tags start and end positions, plus I don't have to manipulate the content. Only the tags

func (l *Lexer) readReference() {
	for isChar(l.char) {
		l.readChar()
	}
	if l.peek() == '>' {
		l.readChar()
	} else {
		log.Panic("Expected >")
	}
}

func (l *Lexer) NextToken() token.Token {

	// current doesn't advance so we know the start position (curr) and end position (curr) of a token

	l.startPos = l.currPos

	// as long start == false keep reading char
	// check encounter <article>
	// set start == true

	// NOTE no pointer because the struct is very small.
	// TODO test impact of using pointer once project is finished
	tok := token.Token{}

	for l.start == false {
		switch l.char {

		case '<':
			//NOTE consume '<'
			l.readChar()
			l.readReference()
			tok = l.makeToken()
			if tok.Type == token.OPEN_ARTICLE {
				l.start = true
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
			break
		case '/':
			fmt.Println("/ found")
			l.readChar()
			break
		case '>':
			fmt.Println("> found. Make token")
			l.readChar()
			break
		default:
			if isChar(l.char) {
				fmt.Printf("char found %c \n", l.char)
				//NOTE I prefer to not call readchar in the isChar funciton so it's more explicit as to whereI'm call the is readchar function
				l.readChar()
				break
			}
			fmt.Println("no char found")
			break

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
