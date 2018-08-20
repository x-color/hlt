package main

import (
	"errors"
	"strings"
	"testing"
)

func TestGenCharColor(t *testing.T) {
	expected := "\x1b[38;5;9m"
	actual := genCharColor(9)
	if actual != expected {
		msg := "Didn't generate correct charactor color code"
		t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
	}
}

func TestGenBackColor(t *testing.T) {
	expected := "\x1b[48;5;9m"
	actual := genBackColor(9)
	if actual != expected {
		msg := "Didn't generate correct background color code"
		t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
	}
}

func TestGenStyleCode(t *testing.T) {
	testCase := map[string][]int{
		"\x1b[38;5;0m\x1b[48;5;255m": []int{0, 255},
		"\x1b[38;5;0m":               []int{0, 256},
		"\x1b[48;5;255m":             []int{-1, 255},
		"":                           []int{256, -1},
	}
	for expected, set := range testCase {
		actual := genStyleCode(set[0], set[1])
		if actual != expected {
			msg := "Didn't generate correct color code"
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, []byte(expected), []byte(actual))
		}
	}
}

func TestHighlightLines(t *testing.T) {
	colorCode := "\x1b[38;5;9m"
	lines := make(chan string, 2)
	output := make(chan string, 2)
	expected := "\x1b[38;5;9mhello world\x1b[0m\ntest message\n"
	actual := ""
	go hightlightLines("world", colorCode, lines, output)
	lines <- "hello world"
	lines <- "test message"
	close(lines)
	for line := range output {
		actual = strings.Join([]string{actual, line, "\n"}, "")
	}
	if actual != expected {
		msg := "Didn't return correct the sets of first index and last index"
		t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
	}
}

func TestHighlightText(t *testing.T) {
	colorCode := "\x1b[38;5;9m"
	lines := make(chan string, 2)
	output := make(chan string, 2)
	expected := "hello \x1b[38;5;9mworld\x1b[0m\ntest message\n"
	actual := ""
	go hightlightText("world", colorCode, lines, output)
	lines <- "hello world"
	lines <- "test message"
	close(lines)
	for line := range output {
		actual = strings.Join([]string{actual, line, "\n"}, "")
	}
	if actual != expected {
		msg := "Didn't return correct the sets of first index and last index"
		t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
	}
}

func TestColorToNumCorrectString(t *testing.T) {
	testCase := map[string]int{
		"red":  196,
		"56":   56,
		"none": -1,
	}
	for input, expected := range testCase {
		actual, err := colorToNum(input)
		if err != nil {
			msg := "Didn't convert to correct color number. Catched error"
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, err)
		}
		if actual != expected {
			msg := "Didn't convert correct color number"
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
		}
	}
}

func TestColorToNumIncorrectString(t *testing.T) {
	testCase := map[string]error{
		"test": errors.New("not defined color"),
		"a56":  errors.New("not defined color"),
	}
	for input, expected := range testCase {
		_, actual := colorToNum(input)
		if actual == nil {
			msg := "Didn't catched error of converting incorrect color"
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
		}
		if actual.Error() != expected.Error() {
			msg := "Didn't convert correct color number"
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
		}
	}
}
