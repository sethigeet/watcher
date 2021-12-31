// Watcher is a command line tool that watches the files in the current
// directory and when any of the files change, it runs the specified command
package main

import (
	"fmt"
	"os"

	"github.com/sethigeet/watcher/watcher"
	"github.com/sethigeet/watcher/watcher/cmd"
)

func main() {
	cmd.Setup()

	exitCode, err := cmd.Parse()
	if err != nil {
		cmd.PrintUsage(err)
		os.Exit(exitCode)
	}

	err = watcher.Setup()
	if err != nil {
		fmt.Printf("An error occurred!\nerror: %s\n", err)
		return
	}
}
