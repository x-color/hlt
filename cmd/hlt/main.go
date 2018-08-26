package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

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
	app.Usage = "highlight texts or lines matched a given pattern in files"
	app.Version = "0.3.1"
	app.Author = "x-color"
	app.HelpName = app.Name
	app.UsageText = app.Name + " [global option] command [option]... [argument]..."
	app.Description = strings.Join([]string{
		app.Name,
		"highlights texts or lines matched a given pattern in files (or standard input if no files set to arguments).",
	}, " ")
	app.HideHelp = true

	// It is flag of highlight commands ('word', 'line')
	highlightFlags := []cli.Flag{
		cli.StringFlag{
			Name:        "background, b",
			Value:       "none",
			Usage:       "`color` background",
			Destination: &opt.background,
		},
		cli.StringFlag{
			Name:        "charactor, c",
			Value:       "red",
			Usage:       "`color` charactor",
			Destination: &opt.charactor,
		},
		cli.BoolFlag{
			Name:        "bold, B",
			Usage:       "bold format",
			Destination: &opt.bold,
		},
		cli.BoolFlag{
			Name:        "hide, H",
			Usage:       "hide text",
			Destination: &opt.hide,
		},
		cli.BoolFlag{
			Name:        "italic, I",
			Usage:       "italic format",
			Destination: &opt.italic,
		},
		cli.BoolFlag{
			Name:        "strikethrough, S",
			Usage:       "strikethrough text",
			Destination: &opt.strikethrough,
		},
		cli.BoolFlag{
			Name:        "underline, U",
			Usage:       "underline text",
			Destination: &opt.underline,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:      "line",
			Aliases:   []string{"l"},
			Usage:     "highlight lines containing a text matched given patten",
			UsageText: app.Name + " line [option]... pattern [file]...",
			Description: strings.Join([]string{
				"It highlights lines containing a text matched given pattern in files (or standard input if no files set to arguments).",
				"It highlights only charactor on by default.",
				"Settable color is 'none','black','blue','cyan','green','magenta','red','yellow' and 0~255.",
				"The color number(0~255) is supported by some terminals.",
			}, " "),
			Flags: append(highlightFlags, []cli.Flag{
				cli.IntFlag{
					Name:        "after",
					Value:       0,
					Usage:       "highlight `num` lines after matching lines",
					Destination: &opt.after,
				},
				cli.IntFlag{
					Name:        "before",
					Value:       0,
					Usage:       "highlight `num` lines before matching lines",
					Destination: &opt.before,
				},
			}...),
			Action: highlightAction(hightlightLines),
		},
		{
			Name:      "word",
			Aliases:   []string{"w"},
			Usage:     "highlight texts matched given patten",
			UsageText: app.Name + " word [option]... pattern [file]...",
			Description: strings.Join([]string{
				"It highlights texts matched given pattern in files (or standard input if no files set to arguments).",
				"It highlights only charactor on by default.",
				"Settable color is 'none','black','blue','cyan','green','magenta','red','yellow' and 0~255.",
				"The color number(0~255) is supported by some terminals.",
			}, " "),
			Flags:  highlightFlags,
			Action: highlightAction(hightlightText),
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
