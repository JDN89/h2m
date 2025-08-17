package token

// NOTE typedefiniton TokenType is other name for int
type TokenType int

type Token struct {
	Type TokenType
	Pos  int
}

const (
	GT TokenType = iota
	LT
	SLASH
	DIV
	OPEN_DIV
	CLOSED_DIV
	OPEN_ARTICLE
	CLOSED_ARTICLE
	EOF
	ERROR_NO_CLOSING_TAG
	CONTENT
)

// NOTE I can leave TokenType out of the onst enum expression and implictly cast the int to TokenType
var HtmlReferenceTokenMap = map[string]TokenType{
	"<article>":  OPEN_ARTICLE,
	"</article>": CLOSED_ARTICLE,
}
