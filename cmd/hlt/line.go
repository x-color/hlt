package main

import "strings"

// LineOpt is option of line command
type LineOpt struct {
	after  int
	before int
}

// LineArg is argument of line command
type LineArg struct {
	pattern string
}

// highlightLines adds color code to head and tail of line including pattern
// and sends it to channel
func highlightLines(colorCode string, lines, output chan string) {
	buffer := []string{}
	after := 0
	for line := range lines {
		if len(colorCode) == 0 {
			output <- line
			continue
		}
		if len(buffer) > opt.line.before {
			output <- buffer[0]
			buffer = buffer[1:]
		}
		buffer = append(buffer, line)
		switch {
		case strings.Contains(line, arg.line.pattern):
			after = opt.line.after + 1
			fallthrough
		case after > 0:
			for _, l := range buffer {
				output <- colorCode + l + "\x1b[0m"
			}
			buffer = buffer[:0]
		}
		after--
	}
	for _, l := range buffer {
		output <- l
	}
	close(output)
}
