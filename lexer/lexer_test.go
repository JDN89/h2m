package lexer
import(	"h2m/token"
    "testing"
)

func TestGetNextToken(t *testing.T) {
	input  := "<div> content </div>";

	expectedTokens  := [2]token.TokenType{token.OPEN_DIV,token.CLOSED_DIV}

	l := New(input)

	for i, tt := range expectedTokens {
		tok := l.nextToken()
		if tok.Type != tt{
			t.Fatalf("Expected token %s but got %s!",expectedTokens[i],tok.Type)
		}
	}
}
