package commands

import "fmt"

func Help() {
	fmt.Println("LocalShare — Temporary file sharing over your LAN.")
	fmt.Println()

	fmt.Println("Usage:")
	fmt.Println("  localshare <command> [flags]")
	fmt.Println()

	fmt.Println("Commands:")
	fmt.Println("  share        Share files with devices on the same local network.")
	fmt.Println("  help         Show help for a command.")
	fmt.Println()

	fmt.Println("Examples:")
	fmt.Println("  localshare share --files photo.png notes.txt")
	fmt.Println("  localshare help")
}

func ShareHelp() {
	fmt.Println("LocalShare — Temporary file sharing over your LAN.")
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
}
