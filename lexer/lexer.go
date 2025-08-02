package lexer

import(	"github.com/JDN89/h2m/token"
)


//TODO make input private. Now public for debug purposes
type Lexer struct {
	Input string
	curr  int
	next  int
	char  byte
}

func New(input string) *Lexer {
	l := &Lexer{Input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.next >= len(l.Input) {
		l.char = 0
	} else {
		l.char = l.Input[l.next]
	}
	l.curr = l.next
	l.next ++
}

//TODO define tokens
// Type and position in input stream
// no need to copy the string, calculate the offset
// identifier tokens <p> </p>
//content Token -> part between <p> </p>
//just track start position  end of contennt is start of enclosing tag

// NOTE: I noticed that the to blogs that I want to convert to markdown the blog content is wrapped in the <article> elmenent. So I'll start lexing once I encounter article. When encounter '<' identify which element it is and store the type and position in a token struct. for the content als just the position is needed. altough why? we just need to replace the html tags and not the content. So I can probably start with just replacing the tags.
//WHAT with tags that immediatley have a closing tag, tables, hrefs, lists? worries for later
