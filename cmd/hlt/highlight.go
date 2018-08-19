package main

import (
	"fmt"
	"strings"
)

// genCharColor generates charactor color code from color number
func genCharColor(num int) (colorCode string) {
	return fmt.Sprintf("\x1b[38;5;%dm", num)
}

// genBackColor generates background color code from color number
func genBackColor(num int) (colorCode string) {
	return fmt.Sprintf("\x1b[48;5;%dm", num)
}

// genColorCode generates color code from color number
func genColorCode(charColor, backColor int) (colorCode string) {
	if 0 <= charColor && charColor <= 255 {
		colorCode = strings.Join([]string{colorCode, genCharColor(charColor)}, "")
	}
	if 0 <= backColor && backColor <= 255 {
		colorCode = strings.Join([]string{colorCode, genBackColor(backColor)}, "")
	}
	return colorCode
}

// hightlightLines adds color code to head and tail of line including pattern
// and sends it to channel
func hightlightLines(pattern, colorCode string, lines, output chan string) {
	for line := range lines {
		if len(colorCode) > 0 && strings.Contains(line, pattern) {
			output <- colorCode + line + "\x1b[0m"
		} else {
			output <- line
		}
	}
	close(output)
}

// hightlightText adds color code to head and tail of text matching a pattern
// and sends it to channel
func hightlightText(pattern, colorCode string, lines, output chan string) {
	for line := range lines {
		if len(colorCode) > 0 && strings.Contains(line, pattern) {
			output <- strings.Replace(line, pattern, colorCode+pattern+"\x1b[0m", -1)
		} else {
			output <- line
		}
	}
	close(output)
}

func hightlightProcess(arg Argument, opt Option, addColor func(string, string, chan string, chan string)) {
	lines := make(chan string, 100)
	output := make(chan string, 100)
	if len(arg.files) == 0 {
		go readStdin(lines)
	} else {
		go readFiles(arg.files, lines)
	}
	colorCode := genColorCode(opt.charactor, opt.background)
	go addColor(arg.pattern, colorCode, lines, output)
	for line := range output {
		fmt.Println(line)
	}
}
