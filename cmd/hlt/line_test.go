package main

import (
	"testing"
)

func TestHighlightLines(t *testing.T) {
	testCases := []struct {
		colorCode string
		after     int
		before    int
		output    string
	}{
		{"\x1b[38;5;9m", 0, 0, "hello world\n\x1b[38;5;9mtest message\x1b[0m\nfinish\n"},
		{"\x1b[38;5;9m", 1, 0, "hello world\n\x1b[38;5;9mtest message\x1b[0m\n\x1b[38;5;9mfinish\x1b[0m\n"},
		{"\x1b[38;5;9m", 0, 1, "\x1b[38;5;9mhello world\x1b[0m\n\x1b[38;5;9mtest message\x1b[0m\nfinish\n"},
		{"\x1b[38;5;9m", 1, 1, "\x1b[38;5;9mhello world\x1b[0m\n\x1b[38;5;9mtest message\x1b[0m\n\x1b[38;5;9mfinish\x1b[0m\n"},
		{"", 0, 0, "hello world\ntest message\nfinish\n"},
	}
	arg.line.pattern = "test"
	for _, tc := range testCases {
		expected := tc.output
		actual := ""
		opt.line.after, opt.line.before = tc.after, tc.before
		lines := make(chan string, 2)
		output := make(chan string, 2)
		go highlightLines(tc.colorCode, lines, output)
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
