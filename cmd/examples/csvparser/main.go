package main

import (
	"regexp"
	"strings"
)

var withinQuotesRegexp = regexp.MustCompile(`(".+")`)

func parseCSVLine(input string) []string {

	matches := withinQuotesRegexp.FindAllString(input, -1)

	// Remove already matched tokens in input.
	modifiedInput := input
	for _, match := range matches {
		modifiedInput = strings.ReplaceAll(modifiedInput, match + ",", "")
	}

	// Splice the fields within quotes, with those which are raw comma separated.
	results := append(matches, strings.Split(modifiedInput, ",")...)
	return results
}

func main() {

}
