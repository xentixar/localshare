package commands

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

var flags = []string{
	"--files",
	"--password",
	"--expire",
	"--port",
}

func Share(args []string) {
	files := []string{}
	password := ""
	expire := 0
	port := 8092

	discoveredFiles := false
	discoveredPassword := false
	discoveredExpire := false
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
		} else if arg == "--expire" && !discoveredExpire {
			expire = castToInt(args[i+1], "expire")
			discoveredExpire = true
		} else if arg == "--port" && !discoveredPort {
			port = castToInt(args[i+1], "port")
			discoveredPort = true
		}
	}

	for _, file := range files {
		checkFileIsvalid(file)
	}

	fmt.Println(files, password, expire, port)
}

func checkFileIsvalid(file string) {
	info, err := os.Stat(file)
	if err != nil {
		fmt.Println("File " + file + " doesn't exists.")
		os.Exit(1)
	}

	if info.IsDir() {
		fmt.Println("Expecting " + file + " a file but it's a folder.")
		os.Exit(1)
	}
}

func castToInt(num string, kind string) int {
	n, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println("Invalid " + kind)
		os.Exit(1)
	}

	return n
}
