package main

import "testing"

func TestInNums(t *testing.T) {
	testCase := []struct {
		num    int
		output bool
	}{
		{1, true},
		{2, false},
		{3, true},
		{4, false},
		{8, true},
		{9, false},
	}
	nums := map[int]bool{
		1: true,
		3: true,
		8: true,
	}
	checker := inNums(nums)
	for _, tc := range testCase {
		expected := tc.output
		actual := checker(tc.num)
		if actual != expected {
			msg := "Didn't check the number correctory"
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
		}
	}
}

func TestIsOdd(t *testing.T) {
	testCase := []struct {
		num    int
		output bool
	}{
		{1, true},
		{2, false},
		{3, true},
		{4, false},
		{8, false},
		{9, true},
	}
	for _, tc := range testCase {
		expected := tc.output
		actual := isOdd(tc.num)
		if actual != expected {
			msg := "Didn't check the odd number"
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
		}
	}
}

func TestIsEven(t *testing.T) {
	testCase := []struct {
		num    int
		output bool
	}{
		{1, false},
		{2, true},
		{3, false},
		{4, true},
		{8, true},
		{9, false},
	}
	for _, tc := range testCase {
		expected := tc.output
		actual := isEven(tc.num)
		if actual != expected {
			msg := "Didn't check the even number"
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
		}
	}
}

func TestIsBetweenAandB(t *testing.T) {
	testCase := []struct {
		num    int
		output bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, true},
		{5, false},
	}
	checker := isBetweenAandB(2, 4)
	for _, tc := range testCase {
		expected := tc.output
		actual := checker(tc.num)
		if actual != expected {
			msg := "Didn't check the even number"
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
		}
	}
}

func TestIsAssignedNum(t *testing.T) {
	testCase := []struct {
		num    int
		output bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, true},
	}
	checkers = append(checkers, isOdd, inNums(map[int]bool{2: true}), isBetweenAandB(8, 10))
	for _, tc := range testCase {
		expected := tc.output
		actual := isAssignedNum(tc.num)
		if actual != expected {
			msg := "Didn't check the number correctory"
			t.Fatalf("%s\nExpected: %v\nActual  : %v", msg, expected, actual)
		}
	}
}

func TestHighlightNumLines(t *testing.T) {
	testCase := []struct {
		checkers []func(int) bool
		output   string
	}{
		{
			[]func(int) bool{isEven},
			"1\n\x1b[38;5;9m2\x1b[0m\n3\n\x1b[38;5;9m4\x1b[0m\n5\n",
		},
		{
			[]func(int) bool{isOdd},
			"\x1b[38;5;9m1\x1b[0m\n2\n\x1b[38;5;9m3\x1b[0m\n4\n\x1b[38;5;9m5\x1b[0m\n",
		},
		{
			[]func(int) bool{inNums(map[int]bool{2: true, 4: true})},
			"1\n\x1b[38;5;9m2\x1b[0m\n3\n\x1b[38;5;9m4\x1b[0m\n5\n",
		},
		{
			[]func(int) bool{isBetweenAandB(2, 4)},
			"1\n\x1b[38;5;9m2\x1b[0m\n\x1b[38;5;9m3\x1b[0m\n\x1b[38;5;9m4\x1b[0m\n5\n",
		},
	}
	colorCode := "\x1b[38;5;9m"
	for _, tc := range testCase {
		expected := tc.output
		actual := ""
		checkers = tc.checkers
		lines := make(chan string, 2)
		output := make(chan string, 2)
		go hightlightNumLines(colorCode, lines, output)
		lines <- "1"
		lines <- "2"
		lines <- "3"
		lines <- "4"
		lines <- "5"
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
