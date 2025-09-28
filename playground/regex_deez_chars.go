package playground

import (
	"fmt"
	"regexp"
)

func ParseHref(input []byte) map[string]string {

	attributes := make(map[string]string)

	var hrefRe = regexp.MustCompile(`href="([^"]+)"`)

	// FindSubmatch returns the whole match and then the capture groups if any are found
	href := hrefRe.FindSubmatch(input)

	if len(href) > 1 {
		attributes["href"] = string(href[1])
	}

	fmt.Println(attributes["href"])

	return attributes

}
