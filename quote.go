// Package shquote helps with quoting strings that will be interpretd by a UNIX
// shell.
package shquote

import (
	"regexp"
	"strings"
)

var unsafe = regexp.MustCompile("[^\\w@%+=:,./-]")

// Quote returns a shell-quoted version of the input string.
func Quote(str string) string {
	// empty string gets special treatment
	if str == "" {
		return "''"
	}

	// if there are no unsafe characters, no quoting is required
	if unsafe.FindStringIndex(str) == nil {
		return str
	}

	// use single quotes, and double-quote embedded single quotes
	return "'" + strings.Replace(str, "'", "'\"'\"'", -1) + "'"
}

// QuoteList quotes each element in the list and joins them with whitespace.
func QuoteList(unquoted []string) string {
	var quoted = make([]string, len(unquoted))
	for i, u := range unquoted {
		quoted[i] = Quote(u)
	}

	return strings.Join(quoted, " ")
}
