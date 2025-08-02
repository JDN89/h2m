package token

type TokenType string;

type Token struct{
TokenType string
pos int
}

const (
	GT = ">";
	LT = "<"
)
