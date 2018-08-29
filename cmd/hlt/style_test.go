package main

import (
	"testing"
)

func TestGenCharColor(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"9", "\x1b[38;5;9m"},
		{"red", "\x1b[31m"},
		{"none", ""},
		{"256", ""},
	}
	for _, tc := range testCases {
		expected := tc.output
		actual := genCharColor(tc.input)
		if actual != expected {
			msg := "Didn't generate correct charactor color code"
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, []byte(expected), []byte(actual))
		}
	}
}

func TestGenBackColor(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"9", "\x1b[48;5;9m"},
		{"red", "\x1b[41m"},
		{"none", ""},
		{"256", ""},
	}
	for _, tc := range testCases {
		expected := tc.output
		actual := genBackColor(tc.input)
		if actual != expected {
			msg := "Didn't generate correct background color code"
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, []byte(expected), []byte(actual))
		}
	}
}

func TestGenBoldStyle(t *testing.T) {
	testCases := []struct {
		input  bool
		output string
	}{
		{true, "\x1b[1m"},
		{false, ""},
	}
	for _, tc := range testCases {
		expected := tc.output
		actual := genBoldStyle(tc.input)
		if actual != expected {
			msg := "Didn't generate correct style code"
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, []byte(expected), []byte(actual))
		}
	}
}

func TestGenHideStyle(t *testing.T) {
	testCases := []struct {
		input  bool
		output string
	}{
		{true, "\x1b[8m"},
		{false, ""},
	}
	for _, tc := range testCases {
		expected := tc.output
		actual := genHideStyle(tc.input)
		if actual != expected {
			msg := "Didn't generate correct style code"
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, []byte(expected), []byte(actual))
		}
	}
}

func TestGenItalicStyle(t *testing.T) {
	testCases := []struct {
		input  bool
		output string
	}{
		{true, "\x1b[3m"},
		{false, ""},
	}
	for _, tc := range testCases {
		expected := tc.output
		actual := genItalicStyle(tc.input)
		if actual != expected {
			msg := "Didn't generate correct style code"
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, []byte(expected), []byte(actual))
		}
	}
}

func TestGenStrikethroughStyle(t *testing.T) {
	testCases := []struct {
		input  bool
		output string
	}{
		{true, "\x1b[9m"},
		{false, ""},
	}
	for _, tc := range testCases {
		expected := tc.output
		actual := genStrikethroughStyle(tc.input)
		if actual != expected {
			msg := "Didn't generate correct style code"
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, []byte(expected), []byte(actual))
		}
	}
}

func TestGenUnderlineStyle(t *testing.T) {
	testCases := []struct {
		input  bool
		output string
	}{
		{true, "\x1b[4m"},
		{false, ""},
	}
	for _, tc := range testCases {
		expected := tc.output
		actual := genUnderlineStyle(tc.input)
		if actual != expected {
			msg := "Didn't generate correct style code"
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, []byte(expected), []byte(actual))
		}
	}
}

func TestGenStyleCode(t *testing.T) {
	opt.style.Charactor = "red"
	opt.style.Background = "red"
	opt.style.Bold = true
	opt.style.Hide = true
	opt.style.Italic = true
	opt.style.Strikethrough = true
	opt.style.Underline = true
	expected := "\x1b[31m\x1b[41m\x1b[1m\x1b[8m\x1b[3m\x1b[9m\x1b[4m"
	actual := genStyleCode()
	if actual != expected {
		msg := "Didn't generate correct style code"
		t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, []byte(expected), []byte(actual))
	}
}
