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
	return
}

func isChar(char byte) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z'
}

func (l *Lexer) NextToken() token.Token {
	// current doesn't advance so we know the start position (curr) and end position (curr) of a token

	l.startPos = l.currPos

	// as long start == false keep reading char
	// check encounter <article>
	// set start == true
	for l.start == false {
		// TODO stop parsing if we reach end of file and never encounter article

	}

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
	token := token.Token{
		Type: token.LT,
		Pos:  5}
	return token
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
