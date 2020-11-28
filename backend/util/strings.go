package util

import (
	"bufio"
	"strings"
	"unicode"
)

// Abbrev returns an abbreviated version of `s`, no longer than `n` characters long.
// It calls NormalizeWhitespace on the input
func Abbrev(s string, n int) string {
	s = NormalizeWhitespace(s)
	if len(s) > n {
		suffix := "..."
		return s[0:n-len(suffix)] + suffix
	}
	return s
}

// NormalizeWhitespace trims leading and trailing whitespace from each line,
// preserving indentation relative to the first line, and replacing windows-style
// `\r\n` with just `\n`
func NormalizeWhitespace(input string) string {
	output := ""

	// note that a simple `strings.Split(s, "\n")` will fail on windows-style `\r\n` sequences
	// `bufio.NewScanner` avoids that problem.
	scan := bufio.NewScanner(strings.NewReader(input))

	isLeading := true
	indentation := ""

	for scan.Scan() {
		// start by trimming trailing whitespace
		// can't trim leading whitespace yet because we need to measure any indentation
		line := strings.TrimRightFunc(scan.Text(), unicode.IsSpace)

		if isLeading {
			// we don't care at all about empty lines before the first non-whitespace character
			if len(line) == 0 {
				continue
			}

			// find the first non-whitespace character
			if index := strings.IndexFunc(line, notSpace); index > 0 {
				// once we hit our first character, lines are no longer "leading", and we know what our indentation prefix is
				indentation = line[0:index]
				isLeading = false
			}

			output += strings.TrimPrefix(line, indentation)
		} else {
			// all lines after the first should be preceded with a plain old newline
			output += "\n" + strings.TrimPrefix(line, indentation)
		}
	}

	return output
}

func notSpace(r rune) bool {
	return !unicode.IsSpace(r)
}
