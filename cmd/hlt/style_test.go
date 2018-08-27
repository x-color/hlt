package main

import (
	"testing"
)

func TestGenCharColor(t *testing.T) {
	testCase := map[string]string{
		"9":    "\x1b[38;5;9m",
		"red":  "\x1b[31m",
		"none": "",
		"256":  "",
	}
	for color, expected := range testCase {
		actual := genCharColor(color)
		if actual != expected {
			msg := "Didn't generate correct charactor color code"
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, []byte(expected), []byte(actual))
		}
	}
}

func TestGenBackColor(t *testing.T) {
	testCase := map[string]string{
		"9":    "\x1b[48;5;9m",
		"red":  "\x1b[41m",
		"none": "",
		"256":  "",
	}
	for color, expected := range testCase {
		actual := genBackColor(color)
		if actual != expected {
			msg := "Didn't generate correct background color code"
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, []byte(expected), []byte(actual))
		}
	}
}

func TestGenBoldStyle(t *testing.T) {
	expected := "\x1b[1m"
	actual := genBoldStyle(true)
	if actual != expected {
		msg := "Didn't generate correct style code"
		t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, []byte(expected), []byte(actual))
	}
}

func TestGenHideStyle(t *testing.T) {
	expected := "\x1b[8m"
	actual := genHideStyle(true)
	if actual != expected {
		msg := "Didn't generate correct style code"
		t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, []byte(expected), []byte(actual))
	}
}

func TestGenItalicStyle(t *testing.T) {
	expected := "\x1b[3m"
	actual := genItalicStyle(true)
	if actual != expected {
		msg := "Didn't generate correct style code"
		t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, []byte(expected), []byte(actual))
	}
}

func TestGenStrikethroughStyle(t *testing.T) {
	expected := "\x1b[9m"
	actual := genStrikethroughStyle(true)
	if actual != expected {
		msg := "Didn't generate correct style code"
		t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, []byte(expected), []byte(actual))
	}
}

func TestGenUnderlineStyle(t *testing.T) {
	expected := "\x1b[4m"
	actual := genUnderlineStyle(true)
	if actual != expected {
		msg := "Didn't generate correct style code"
		t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, []byte(expected), []byte(actual))
	}
}
