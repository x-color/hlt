package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

// Option sets options of commands
type Option struct {
	line  LineOpt
	style Style
}

// Argument sets arguments of commands
type Argument struct {
	line  LineArg
	word  WordArg
	lineN lineNArg
	files []string
}

var arg = Argument{}
var opt = Option{}

func usageError(name, usage, message string) {
	fmt.Fprintln(os.Stderr, "Incorrect Usage.", message)
	fmt.Fprintln(os.Stderr, "Usage:", usage)
	fmt.Fprintf(os.Stderr, "If you want more information, execute '%s --help'", name)
}

func main() {
	// Set application's infomations
	app := cli.NewApp()
	app.Name = "hlt"
	app.Usage = "highlight texts or lines in files"
	app.Version = "0.4.5"
	app.Author = "x-color"
	app.HelpName = app.Name
	app.UsageText = app.Name + " [global option] command [option]... [argument]..."
	app.Description = strings.Join([]string{
		app.Name,
		"highlights texts or lines in files (or standard input if no files set to arguments).",
	}, " ")
	app.HideHelp = true

	// It is flag of highlight commands
	highlightFlags := []cli.Flag{
		cli.StringFlag{
			Name:        "background, b",
			Value:       "none",
			Usage:       "`color` background",
			Destination: &opt.style.background,
		},
		cli.StringFlag{
			Name:        "charactor, c",
			Value:       "red",
			Usage:       "`color` charactor",
			Destination: &opt.style.charactor,
		},
		cli.BoolFlag{
			Name:        "bold, B",
			Usage:       "bold format",
			Destination: &opt.style.bold,
		},
		cli.BoolFlag{
			Name:        "hide, H",
			Usage:       "hide text",
			Destination: &opt.style.hide,
		},
		cli.BoolFlag{
			Name:        "italic, I",
			Usage:       "italic format",
			Destination: &opt.style.italic,
		},
		cli.BoolFlag{
			Name:        "strikethrough, S",
			Usage:       "strikethrough text",
			Destination: &opt.style.strikethrough,
		},
		cli.BoolFlag{
			Name:        "underline, U",
			Usage:       "underline text",
			Destination: &opt.style.underline,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:      "line",
			Aliases:   []string{"l"},
			Usage:     "highlight lines containing a text matched given pattern",
			UsageText: app.Name + " line [option]... pattern [file]...",
			Description: strings.Join([]string{
				"It highlights lines containing a text matched given pattern in files (or standard input if no files set to arguments). It highlights only charactor on by default.",
				"",
				"   Use -b, -c to assign highlight color. Settable color is",
				"     - name   'none','black','blue','cyan','green','magenta','red','yellow'",
				"     - number color number is 0~255, it is supported by some terminals",
			}, "\n"),
			Flags: append(highlightFlags, []cli.Flag{
				cli.IntFlag{
					Name:        "after",
					Value:       0,
					Usage:       "highlight `num` lines after matching lines",
					Destination: &opt.line.after,
				},
				cli.IntFlag{
					Name:        "before",
					Value:       0,
					Usage:       "highlight `num` lines before matching lines",
					Destination: &opt.line.before,
				},
			}...),
			Action: func(c *cli.Context) {
				switch {
				case c.NArg() >= 2:
					arg.files = c.Args()[1:]
					fallthrough
				case c.NArg() == 1:
					arg.line.pattern = c.Args()[0]
					arg.word.pattern = c.Args()[0]
				default:
					usageError(app.Name, app.UsageText, "no arguments")
					return
				}
				highlightProcess(highlightLines)
			},
		},
		{
			Name:      "word",
			Aliases:   []string{"w"},
			Usage:     "highlight texts matched given pattern",
			UsageText: app.Name + " word [option]... pattern [file]...",
			Description: strings.Join([]string{
				"It highlights texts matched given pattern in files (or standard input if no files set to arguments). It highlights only charactor on by default.",
				"",
				"   Use -b, -c to assign highlight color. Settable color is",
				"     - name   'none','black','blue','cyan','green','magenta','red','yellow'",
				"     - number color number is 0~255, it is supported by some terminals",
			}, "\n"),
			Flags: highlightFlags,
			Action: func(c *cli.Context) {
				switch {
				case c.NArg() >= 2:
					arg.files = c.Args()[1:]
					fallthrough
				case c.NArg() == 1:
					arg.line.pattern = c.Args()[0]
					arg.word.pattern = c.Args()[0]
				default:
					usageError(app.Name, app.UsageText, "no arguments")
					return
				}
				highlightProcess(highlightText)
			},
		},
		{
			Name:      "linen",
			Aliases:   []string{"ln"},
			Usage:     "highlight assigned lines",
			UsageText: app.Name + " linen [option]... list [file]...",
			Description: strings.Join([]string{
				"It highlights lines assigned by list in files (or standard input if no files set to arguments). It highlights only charactor on by default.",
				"",
				"   list is made up of an integer, a type of number, a range of integers, or multiple integers ranges separated by commas. list consists of",
				"     - N    line N counted from 1",
				"     - N~   from line N to the end of the file",
				"     - ~M   from the first of the file to line M",
				"     - N~M  from line N to line M",
				"     - even even-numbered lines",
				"     - odd  odd-numbered lines",
				"",
				"   Use -b, -c to assign highlight color. Settable color is",
				"     - name   'none','black','blue','cyan','green','magenta','red','yellow'",
				"     - number color number is 0~255, it is supported by some terminals",
			}, "\n"),
			Flags: highlightFlags,
			Action: func(c *cli.Context) {
				switch {
				case c.NArg() >= 2:
					arg.files = c.Args()[1:]
					fallthrough
				case c.NArg() == 1:
					arg.lineN.checkers = parseStringOfList(c.Args()[0])
				default:
					usageError(app.Name, app.UsageText, "no argument")
					return
				}
				highlightProcess(highlightNumLines)
			},
		},
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "help, h",
			Usage: "show help",
		},
	}

	app.Action = func(c *cli.Context) {
		if c.Bool("help") {
			cli.ShowAppHelp(c)
			return
		}
		if c.NArg() > 0 {
			usageError(c.App.Name, c.App.UsageText, "invalid command")
		} else {
			usageError(c.App.Name, c.App.UsageText, "no command")
		}
	}

	app.Run(os.Args)
}
