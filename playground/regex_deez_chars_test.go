package playground

import (
	"testing"
)

func TestRegex(t *testing.T) {
	input := `    <a href="/2015/02/01/what-color-is-your-function/" rel="bookmark" title="Permanent Link to What Color is Your Function?">`

	result := ParseHref([]byte(input))

	if result["href"] != "/2015/02/01/what-color-is-your-function/" {
		t.Errorf("expected href to be extracted, got %q", result["href"])
	}
}

