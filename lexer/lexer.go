package lexer

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
