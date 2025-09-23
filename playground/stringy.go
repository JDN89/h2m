package playground

import (
	"fmt"
	"strings"
)

func StringsAreAwesome(testString string) {
	result := strings.Split(testString, `"`)
	fmt.Println(result)

	if strings.Contains(testString, "href") {
		hrefString := strings.TrimLeft(testString, "href")
		hrefSplittedString := strings.Split(hrefString, `"`)
		hrefContent := hrefSplittedString[1]
		fmt.Println(hrefContent)

	}

	length := len(result)

	for i := 0; i < length; i++ {
		fmt.Println(result[i])
	}

	fmt.Println("done")

}
