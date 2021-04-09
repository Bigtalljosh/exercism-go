package acronym

import (
	"regexp"
	"strings"
)

func Abbreviate(s string) string {
	reg, _ := regexp.Compile("[-,.:;_ ]+")
	substr := reg.ReplaceAllString(s, " ")
	wordArray := strings.Split(substr, " ")
	acronym := ""

	for _, value := range wordArray {
		acronym += string(value[0])
	}

	return strings.ToUpper(acronym)
}
