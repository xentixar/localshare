package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"

	"xentixar/localshare/internals/http"
)

var flags = []string{
	"--files",
	"--password",
	"--port",
}

func Share(args []string) {
	files := []string{}
	password := ""
	port := 8092

	discoveredFiles := false
	discoveredPassword := false
	discoveredPort := false

	for i, arg := range args {
		if arg == "--files" && !discoveredFiles {
			for _, file := range args[i+1:] {
				if slices.Contains(flags, file) {
					break
				}
				files = append(files, file)
			}
			discoveredFiles = true
		} else if arg == "--password" && !discoveredPassword {
			password = args[i+1]
			discoveredPassword = true
		} else if arg == "--port" && !discoveredPort {
			port = castToInt(args[i+1], "port")
			discoveredPort = true
		}
	}

	for i, file := range files {
		files[i] = checkFileIsvalid(file)
	}

	http.Start(port, password, files)
}

func checkFileIsvalid(file string) string {
	absPath, err := filepath.Abs(file)
	if err != nil {
		fmt.Printf("Failed to resolve path %q: %v\n", file, err)
		os.Exit(1)
	}

	info, err := os.Stat(absPath)
	if err != nil {
		fmt.Println("File " + file + " doesn't exists.")
		os.Exit(1)
	}

	if info.IsDir() {
		fmt.Println("Expecting " + file + " a file but it's a folder.")
		os.Exit(1)
	}

	return absPath
}

func castToInt(num string, kind string) int {
	n, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("Invalid " + kind)
		os.Exit(1)
	}

	return n
}
