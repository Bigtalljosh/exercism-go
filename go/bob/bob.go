// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	hasLetters, _ := regexp.MatchString("[a-zA-Z]", remark)
	isQuestion := strings.HasSuffix(remark, "?")
	isYelling := strings.ToUpper(remark) == remark

	switch {
	case hasLetters && isQuestion && isYelling:
		return "Calm down, I know what I'm doing!"
	case hasLetters && isYelling:
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	case len(remark) == 0:
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}

}
