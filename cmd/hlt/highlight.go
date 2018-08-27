package main

import (
	"fmt"
)

func highlightProcess(addColor func(string, chan string, chan string)) {
	lines := make(chan string, 100)
	output := make(chan string, 100)
	if len(arg.files) == 0 {
		go readStdin(lines)
	} else {
		go readFiles(arg.files, lines)
	}
	colorCode := genStyleCode()
	go addColor(colorCode, lines, output)
	for line := range output {
		fmt.Println(line)
	}
}
