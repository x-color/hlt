package main

import (
	"bufio"
	"fmt"
	"os"
)

// readStdin reads input data from standard input and sends it to channel
func readStdin(lines chan string) {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		lines <- s.Text()
	}
	close(lines)
}

// readFiles reads input data from files and sends it to channel
func readFiles(files []string, lines chan string) {
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		s := bufio.NewScanner(f)
		for s.Scan() {
			lines <- s.Text()
		}
		if s.Err() != nil {
			continue
		}
	}
	close(lines)
}
