package playground

import (
	"testing"
)

func TestRegex(t *testing.T) {
	input := `    <a href="/2015/02/01/what-color-is-your-function/" rel="bookmark" title="Permanent Link to What Color is Your Function?">`

	result := ParseHref([]byte(input))

	expected := map[string]string{
		"href":  "/2015/02/01/what-color-is-your-function/",
		"rel":   "bookmark",
		"title": "Permanent Link to What Color is Your Function?",
	}

	if len(result) != len(expected) {
		t.Fatalf("number of elements in attributes map doesn't match. We expected %d , but got %d", len(expected), len(result))
	}

	for attr, content := range expected {
		if result[attr] != content {
			t.Fatalf("Expected %s, but got %s ", content, result[attr])
		}
	}

}
