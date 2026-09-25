// Package cli handles all the command line stuffs and method calling
package cli

import (
	"os"
	"slices"

	"xentixar/localshare/internals/cli/commands"
)

var availableCommands = []string{
	"help",
	"share",
}

func Execute() {
	args := os.Args[1:]
	parse(args)
}

func parse(args []string) {
	if len(args) == 0 || args[0] == "help" || !slices.Contains(availableCommands, args[0]) {
		commands.Help()
	} else if args[0] == "share" {
		commands.Share()
	}
}
