package main

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/urfave/cli"
)

// checkers is list of functions check number of line in assigned range
var checkers []func(int) bool

// inNums returns true if argument in nums
func inNums(nums map[int]bool) func(int) bool {
	return func(i int) (ok bool) {
		if _, ok := nums[i]; ok {
			return true
		}
		return false
	}
}

// isOdd returns true if argument is odd number
func isOdd(i int) (ok bool) {
	if i%2 == 1 {
		return true
	}
	return false
}

// isEven returns true if argument is even number
func isEven(i int) (ok bool) {
	if i%2 == 0 {
		return true
	}
	return false
}

// isBetweenAandB returns function of check argument is between a and b
func isBetweenAandB(a, b int) (checker func(int) bool) {
	return func(i int) (ok bool) {
		if a <= i && i <= b {
			return true
		}
		return false
	}
}

// isAssignedNum checks argument to use checker functions
func isAssignedNum(i int) (ok bool) {
	for _, cheker := range checkers {
		if cheker(i) {
			return true
		}
	}
	return false
}

// hightlightNumLines adds color code to head and tail of N lines and sends it
// to channel
func hightlightNumLines(colorCode string, lines, output chan string) {
	i := 0
	for line := range lines {
		i++
		if len(colorCode) > 0 && isAssignedNum(i) {
			output <- colorCode + line + "\x1b[0m"
		} else {
			output <- line
		}
	}
	close(output)
}

// parseLinenArgs is parser of list of ranges
func parseStringOfList(list string) {
	nums := map[int]bool{}
	for _, condition := range strings.Split(list, ",") {
		switch condition {
		case "odd":
			checkers = append(checkers, isOdd)
			continue
		case "even":
			checkers = append(checkers, isEven)
			continue
		}

		num, err := strconv.Atoi(condition)
		if err == nil {
			nums[num] = true
			continue
		}

		ranges := strings.Split(condition, "~")
		if len(ranges) > 1 {
			min, err := strconv.Atoi(ranges[0])
			if err != nil {
				min = 0
			}
			max, err := strconv.Atoi(ranges[len(ranges)-1])
			if err != nil {
				max = math.MaxInt64
			}
			checkers = append(checkers, isBetweenAandB(min, max))
		}
	}
	checkers = append(checkers, inNums(nums))
}

// parseLinenArgs is parser of arguments of linen command
func parseLinenArgs(c *cli.Context) (err error) {
	switch {
	case c.NArg() >= 2:
		arg.files = c.Args()[1:]
		fallthrough
	case c.NArg() == 1:
		parseStringOfList(c.Args()[0])
		return nil
	default:
		return errors.New("no arguments")
	}
}

// linenAction is action of highlight commands
func linenAction() (action func(*cli.Context)) {
	return func(c *cli.Context) {
		err := parseLinenArgs(c)
		if err != nil {
			usageError(c.App.Name, c.App.UsageText, err.Error())
			return
		}
		hightlightProcess(hightlightNumLines)
	}
}
