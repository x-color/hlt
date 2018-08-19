package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

type Option struct {
	background int
	charactor  int
}

type Argument struct {
	pattern string
	files   []string
}

var opt = Option{}
var arg = Argument{}

func highlightAction(addColor func(string, string, chan string, chan string)) (action func(*cli.Context)) {
	return func(c *cli.Context) {
		switch {
		case c.NArg() >= 2:
			arg.files = c.Args()[1:]
			fallthrough
		case c.NArg() == 1:
			arg.pattern = c.Args()[0]
		default:
			fmt.Fprintln(os.Stderr, "Incorrect Usage. no arguments")
			cli.ShowAppHelp(c)
			return
		}
		hightlightProcess(arg, opt, addColor)
	}
}

func main() {
	// Set application's infomations
	app := cli.NewApp()
	app.Name = "hlt"
	app.Usage = "highlight texts or lines matched a given pattern in files"
	app.Version = "0.1.0"
	app.Author = "x-color"
	app.HelpName = app.Name
	app.UsageText = app.Name + " [global option] command [option]... [argument]..."
	app.Description = strings.Join([]string{
		app.Name,
		"highlights texts or lines matched a given pattern in files (or standard input if no files set to arguments).",
	}, " ")
	app.HideHelp = true

	app.Commands = []cli.Command{
		{
			Name:      "line",
			Aliases:   []string{"l"},
			Usage:     "highlight lines containing a text matched given patten",
			UsageText: app.Name + " line [option]... pattern [file]...",
			Description: strings.Join([]string{
				"It highlights lines containing a text matched given pattern in files (or standard input if no files set to arguments).",
				"It highlights only charactor on by default.",
			}, " "),
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:        "background, b",
					Value:       -1,
					Usage:       "color background. `color` is 0~255",
					Destination: &opt.background,
				},
				cli.IntFlag{
					Name:        "charactor, c",
					Value:       1,
					Usage:       "color charactor. `color` is 0~255",
					Destination: &opt.charactor,
				},
			},
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
			}, " "),
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:        "background, b",
					Value:       -1,
					Usage:       "color background. `color` is 0~255",
					Destination: &opt.background,
				},
				cli.IntFlag{
					Name:        "charactor, c",
					Value:       1,
					Usage:       "color charactor. `color` is 0~255",
					Destination: &opt.charactor,
				},
			},
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
			fmt.Fprintln(os.Stderr, "Incorrect Usage. invalid command")
		} else {
			fmt.Fprintln(os.Stderr, "Incorrect Usage. no command")
		}
		fmt.Fprintln(os.Stderr, "Usage:", app.UsageText)
		fmt.Fprintf(os.Stderr, "If you want more information, execute '%s --help'", app.Name)
	}

	app.Run(os.Args)
}
