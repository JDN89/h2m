package lexer

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
