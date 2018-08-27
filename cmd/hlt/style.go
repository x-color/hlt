package main

import (
	"fmt"
	"strconv"
)

// Style is option of highlight style
type Style struct {
	background    string
	charactor     string
	bold          bool
	hide          bool
	italic        bool
	strikethrough bool
	underline     bool
}

// colorNumber maps color name to number
var colorNumber = map[string]int{
	"black":   0,
	"blue":    4,
	"cyan":    6,
	"green":   2,
	"magenta": 5,
	"red":     1,
	"yellow":  3,
}

// genCharColor generates charactor color code from color number
func genCharColor(color string) (colorCode string) {
	// If color is number, it returns string of 256 colors.
	// The string is suppurted by some terminals.
	num, err := strconv.Atoi(color)
	if err == nil {
		if 0 <= num && num <= 255 {
			return fmt.Sprintf("\x1b[38;5;%dm", num)
		}
		return ""
	}
	// If color is color name, it returns 8 colors.
	// The string is suppurted by most terminals.
	num, ok := colorNumber[color]
	if ok {
		return fmt.Sprintf("\x1b[3%dm", num)
	}
	return ""
}

// genBackColor generates background color code from color number
func genBackColor(color string) (colorCode string) {
	// If color is number, it returns string of 256 colors.
	// The string is suppurted by some terminals.
	num, err := strconv.Atoi(color)
	if err == nil {
		if 0 <= num && num <= 255 {
			return fmt.Sprintf("\x1b[48;5;%dm", num)
		}
		return ""
	}
	// If color is color name, it returns 8 colors.
	// The string is suppurted by most terminals.
	num, ok := colorNumber[color]
	if ok {
		return fmt.Sprintf("\x1b[4%dm", num)
	}
	return ""
}

// genBoldStyle generates bold style code
func genBoldStyle(yes bool) (styleCode string) {
	if yes {
		return "\x1b[1m"
	}
	return ""
}

// genHideStyle generates hide style code
func genHideStyle(yes bool) (styleCode string) {
	if yes {
		return "\x1b[8m"
	}
	return ""
}

// genItalicStyle generates italic style code
func genItalicStyle(yes bool) (styleCode string) {
	if yes {
		return "\x1b[3m"
	}
	return ""
}

// genStrikethroughStyle generates strikethrough style code
func genStrikethroughStyle(yes bool) (styleCode string) {
	if yes {
		return "\x1b[9m"
	}
	return ""
}

// genUnderlineStyle generates underline style code
func genUnderlineStyle(yes bool) (styleCode string) {
	if yes {
		return "\x1b[4m"
	}
	return ""
}

// genStyleCode generates color code from color number
func genStyleCode() (colorCode string) {
	colorCode += genCharColor(opt.style.charactor)
	colorCode += genBackColor(opt.style.background)
	colorCode += genBoldStyle(opt.style.bold)
	colorCode += genHideStyle(opt.style.hide)
	colorCode += genItalicStyle(opt.style.italic)
	colorCode += genStrikethroughStyle(opt.style.strikethrough)
	colorCode += genUnderlineStyle(opt.style.underline)
	return colorCode
}
