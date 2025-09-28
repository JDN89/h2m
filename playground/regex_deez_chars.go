package playground

import (
	"regexp"
)

func ParseHref(input []byte) map[string]string {

	hrefRegexPatterns := map[string]*regexp.Regexp{

		"href":  regexp.MustCompile(`href="([^"]+)"`),
		"rel":   regexp.MustCompile(`rel="([^"]+)"`),
		"title": regexp.MustCompile(`title="([^"]+)"`),
	}

	hrefAttributes := make(map[string]string)

	// FindSubmatch returns the whole match and then the capture groups if any are found

	for key, regex := range hrefRegexPatterns {
		attributeMatch := regex.FindSubmatch(input)
		if len(attributeMatch) > 1 {
			hrefAttributes[key] = string(attributeMatch[1])
		}

	}

	return hrefAttributes

}
