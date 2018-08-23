package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

// Option sets options of highlight commands
type Option struct {
	background    string
	charactor     string
	bold          bool
	hide          bool
	italic        bool
	strikethrough bool
	underline     bool
	after         int
	before        int
}

// Argument sets arguments of highlight commands
type Argument struct {
	pattern string
	files   []string
}

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
func genStyleCode(opt Option) (colorCode string) {
	colorCode += genCharColor(opt.charactor)
	colorCode += genBackColor(opt.background)
	colorCode += genBoldStyle(opt.bold)
	colorCode += genHideStyle(opt.hide)
	colorCode += genItalicStyle(opt.italic)
	colorCode += genStrikethroughStyle(opt.strikethrough)
	colorCode += genUnderlineStyle(opt.underline)
	return colorCode
}

// hightlightLines adds color code to head and tail of line including pattern
// and sends it to channel
func hightlightLines(pattern, colorCode string, lines, output chan string) {
	buffer := []string{}
	after := 0
	for line := range lines {
		if len(colorCode) == 0 {
			output <- line
			continue
		}
		if len(buffer) > opt.before {
			output <- buffer[0]
			buffer = buffer[1:]
		}
		buffer = append(buffer, line)
		switch {
		case strings.Contains(line, pattern):
			after = opt.after + 1
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

func hightlightProcess(addColor func(string, string, chan string, chan string)) {
	lines := make(chan string, 100)
	output := make(chan string, 100)
	if len(arg.files) == 0 {
		go readStdin(lines)
	} else {
		go readFiles(arg.files, lines)
	}
	colorCode := genStyleCode(opt)
	go addColor(arg.pattern, colorCode, lines, output)
	for line := range output {
		fmt.Println(line)
	}
}

// setArguments sets arguments of highlight commands to Argument variable
func setArguments(c *cli.Context) (err error) {
	switch {
	case c.NArg() >= 2:
		arg.files = c.Args()[1:]
		fallthrough
	case c.NArg() == 1:
		arg.pattern = c.Args()[0]
	default:
		return errors.New("no arguments")
	}
	return nil
}

// highlightAction is action of highlight commands
func highlightAction(addColor func(string, string, chan string, chan string)) (action func(*cli.Context)) {
	return func(c *cli.Context) {
		err := setArguments(c)
		if err != nil {
			usageError(c.App.Name, c.App.UsageText, err.Error())
			return
		}
		hightlightProcess(addColor)
	}
}
