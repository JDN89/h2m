package converter

import (
	"h2m/token"
)

var HtmlTokenTypeToMarkdownEquivalentMap = map[token.TokenType]string{
	token.OPEN_HEADER:      "",
	token.CLOSED_HEADER:    "",
	token.OPEN_HEADING_1:   "# ",
	token.CONTENT:          "",
	token.CLOSED_HEADING_1: "\n",
	token.EOF:              "",
	token.OPEN_ARTICLE:     "",
	token.CLOSED_ARTICLE:   "",
	token.ANCHOR_CLOSED:    "",
}

func GetMarkdown(t token.TokenType) string {
	if val, ok := HtmlTokenTypeToMarkdownEquivalentMap[t]; ok {
		return val
	} else {
		// fmt.Printf("No markdown equivalent defined for %s \n", token.TokenType.ToString(t))
		return "No markdown equivalent defined for -- " + token.TokenType.ToString(t) + " \n"
	}
}
