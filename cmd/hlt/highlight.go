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
	background int
	charactor  int
}

// Argument sets arguments of highlight commands
type Argument struct {
	pattern string
	files   []string
}

var colorNumber = map[string]int{
	"blue":   21,
	"green":  40,
	"orange": 202,
	"pink":   207,
	"purple": 164,
	"red":    196,
	"yellow": 226,
}

// genCharColor generates charactor color code from color number
func genCharColor(num int) (colorCode string) {
	return fmt.Sprintf("\x1b[38;5;%dm", num)
}

// genBackColor generates background color code from color number
func genBackColor(num int) (colorCode string) {
	return fmt.Sprintf("\x1b[48;5;%dm", num)
}

// genStyleCode generates color code from color number
func genStyleCode(charColor, backColor int) (colorCode string) {
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
	colorCode := genStyleCode(opt.charactor, opt.background)
	go addColor(arg.pattern, colorCode, lines, output)
	for line := range output {
		fmt.Println(line)
	}
}

// colorToNum converts string of color to color number
func colorToNum(color string) (num int, err error) {
	if color == "none" {
		return -1, nil
	}
	num, err = strconv.Atoi(color)
	if err == nil {
		return num, nil
	}
	num, ok := colorNumber[color]
	if !ok {
		return num, errors.New("not defined color")
	}
	return num, nil
}

// setArguments sets arguments of highlight commands to Argument variable
func setArguments(c *cli.Context) (arg Argument, err error) {
	switch {
	case c.NArg() >= 2:
		arg.files = c.Args()[1:]
		fallthrough
	case c.NArg() == 1:
		arg.pattern = c.Args()[0]
	default:
		return arg, errors.New("no arguments")
	}
	return arg, nil
}

// setArguments sets options of highlight commands to Option variable
func setOptions(c *cli.Context) (opt Option, err error) {
	opt.charactor, err = colorToNum(c.String("charactor"))
	if err != nil {
		return opt, err
	}
	opt.background, err = colorToNum(c.String("background"))
	if err != nil {
		usageError(c.App.Name, c.App.UsageText, err.Error())
	}
	return opt, nil
}

// highlightAction is action of highlight commands
func highlightAction(addColor func(string, string, chan string, chan string)) (action func(*cli.Context)) {
	return func(c *cli.Context) {
		arg, err := setArguments(c)
		if err != nil {
			usageError(c.App.Name, c.App.UsageText, err.Error())
			return
		}
		opt, err := setOptions(c)
		if err != nil {
			usageError(c.App.Name, c.App.UsageText, err.Error())
			return
		}
		hightlightProcess(arg, opt, addColor)
	}
}
