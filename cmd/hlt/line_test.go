package main

import (
	"testing"
)

func TestHighlightLines(t *testing.T) {
	testCase := []struct {
		after  int
		before int
		output string
	}{
		{0, 0, "hello world\n\x1b[38;5;9mtest message\x1b[0m\nfinish\n"},
		{1, 0, "hello world\n\x1b[38;5;9mtest message\x1b[0m\n\x1b[38;5;9mfinish\x1b[0m\n"},
		{0, 1, "\x1b[38;5;9mhello world\x1b[0m\n\x1b[38;5;9mtest message\x1b[0m\nfinish\n"},
		{1, 1, "\x1b[38;5;9mhello world\x1b[0m\n\x1b[38;5;9mtest message\x1b[0m\n\x1b[38;5;9mfinish\x1b[0m\n"},
	}
	colorCode := "\x1b[38;5;9m"
	arg.line.pattern = "test"
	for _, tc := range testCase {
		expected := tc.output
		actual := ""
		opt.line.after, opt.line.before = tc.after, tc.before
		lines := make(chan string, 2)
		output := make(chan string, 2)
		go highlightLines(colorCode, lines, output)
		lines <- "hello world"
		lines <- "test message"
		lines <- "finish"
		close(lines)
		for line := range output {
			actual += line + "\n"
		}
		if actual != expected {
			msg := "Didn't return correct the sets of first index and last index"
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
		}
	}
}
