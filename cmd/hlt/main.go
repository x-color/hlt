package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

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
	app.Version = "0.2.0"
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
			Name:  "background, b",
			Value: "none",
			Usage: "`color` background",
		},
		cli.StringFlag{
			Name:  "charactor, c",
			Value: "red",
			Usage: "`color` charactor",
		},
		cli.BoolFlag{
			Name:  "bold, B",
			Usage: "bold format",
		},
		cli.BoolFlag{
			Name:  "hide, H",
			Usage: "hide text",
		},
		cli.BoolFlag{
			Name:  "italic, I",
			Usage: "italic format",
		},
		cli.BoolFlag{
			Name:  "strikethrough, S",
			Usage: "strikethrough text",
		},
		cli.BoolFlag{
			Name:  "underline, U",
			Usage: "underline text",
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
				"Settable color is 0~255 and 'none','blue','green','orange','pink','purple','red','yellow'.",
			}, " "),
			Flags:  highlightFlags,
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
				"Settable color is 0~255 and 'none','blue','green','orange','pink','purple','red','yellow'.",
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
