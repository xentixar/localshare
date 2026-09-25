package commands

import "fmt"

func Help() {
	fmt.Println("Temporary file sharing over your LAN")
	fmt.Println()
	fmt.Println("USAGE")
	fmt.Println("  localshare <command> <subcommand>")
	fmt.Println()
	fmt.Println("COMMANDS")
	fmt.Println("  share:       Share files")
	fmt.Println("  help:        Shows help")
}

func ShareHelp() {
	fmt.Println("Temporary file sharing over your LAN")
	fmt.Println()

	fmt.Println("Usage:")
	fmt.Println("  localshare share [flags]")
	fmt.Println()

	fmt.Println("Flags:")
	fmt.Println("  --files <paths...>")
	fmt.Println("      Files to share.")
	fmt.Println("      Accepts multiple files.")
	fmt.Println()

	fmt.Println("  --password <password>")
	fmt.Println("      Protect the shared files with a password.")
	fmt.Println()

	fmt.Println("  --port <port>")
	fmt.Println("      Port to host the sharing server on.")
	fmt.Println("      Default: 8092")
	fmt.Println()

	fmt.Println("  --expire <seconds>")
	fmt.Println("      Automatically stop sharing after the given number of seconds.")
	fmt.Println("      Default: never expires")
	fmt.Println()
}
