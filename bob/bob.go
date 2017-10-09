package bob

import (
	"regexp"
	"strings"
)

const testVersion = 3

func Hey(msg string) string {
	msg = strings.TrimSpace(msg)

	// address without saying
	if msg == "" {
		return "Fine. Be that way!"
	}

	// yell
	if isYell(msg) {
		return "Whoa, chill out!"
	}

	// question
	if msg[len(msg)-1:] == "?" {
		return "Sure."
	}

	return "Whatever."
}

var hasCharacterRegexp = regexp.MustCompile(`[a-zA-Z]+`)

func isYell(msg string) bool {
	hasCharacter := hasCharacterRegexp.MatchString(msg)
	allUpper := strings.ToUpper(msg) == msg
	return hasCharacter && allUpper
}
