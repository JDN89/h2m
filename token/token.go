package token

type TokenType string

type Token struct {
	Type TokenType
	Pos  int
}

const (
	GT         = TokenType(">")
	LT         = TokenType("<")
	SLASH      = TokenType("/")
	DIV        = TokenType("div")
	OPEN_DIV   = TokenType("<div>")
	CLOSED_DIV = TokenType("</div>")
)
