package main

import (
	"os"
	"strings"
)

var running = true

func main() {
	args := os.Args

	if len(args) < 2 {
		printUsage()
		return
	}

	_, command := parseArguments(args)

	switch command[0] {
	case "help":
		printHelp()
	case "ls", "list":
		prettyPrintListCards(command)
	case "pls", "plainlist":
		plainPrintListCards(command)
	}

}

func parseArguments(args []string) (string, []string) {
	var options string
	var command []string

	if strings.Contains(args[1], "-") && len(args) > 2 {
		options = args[1]
		command = args[2:]
	} else {
		command = args[1:]
	}

	return options, command
}
