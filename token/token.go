package token

// NOTE typedefiniton TokenType is other name for int
type TokenType int

type Token struct {
	Type     TokenType
	StartPos int
	EndPos   int
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
	OPEN_HEADING_1
	CLOSED_HEADING_1
	OPEN_HEADER
	CLOSED_HEADER
	EOF
	ERROR_NO_CLOSING_TAG
	CONTENT
)

// NOTE I can leave TokenType out of the onst enum expression and implictly cast the int to TokenType
var HtmlReferenceTokenMap = map[string]TokenType{
	"<article>":  OPEN_ARTICLE,
	"</article>": CLOSED_ARTICLE,
	"<header>":   OPEN_HEADER,
	"</header>":  CLOSED_HEADER,
	"<h1>":       OPEN_HEADING_1,
	"</h1>":      CLOSED_HEADING_1,
}

var tokenTypeToString = map[TokenType]string{
	OPEN_ARTICLE:     "<article>",
	CLOSED_ARTICLE:   "</article>",
	OPEN_HEADER:      "<header>",
	CLOSED_HEADER:    "</header>",
	OPEN_HEADING_1:   "<h1>",
	CLOSED_HEADING_1: "</h1>",
}

func (t TokenType) ToString() string {

	if val, ok := tokenTypeToString[t]; ok {

		return val
	}
	return "UNKOWN"
}
