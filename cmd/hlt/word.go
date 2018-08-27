package main

import "strings"

// WordArg is argument of word command
type WordArg struct {
	pattern string
}

// highlightText adds color code to head and tail of text matching a pattern
// and sends it to channel
func highlightText(colorCode string, lines, output chan string) {
	for line := range lines {
		if len(colorCode) > 0 && strings.Contains(line, arg.word.pattern) {
			output <- strings.Replace(line, arg.word.pattern, colorCode+arg.word.pattern+"\x1b[0m", -1)
		} else {
			output <- line
		}
	}
	close(output)
}
