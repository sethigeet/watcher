// Watcher is a command line tool that watches the files in the current
// directory and when any of the files change, it runs the specified command
package main

import (
	"os"

	"github.com/op/go-logging"

	"github.com/sethigeet/watcher/watcher"
	"github.com/sethigeet/watcher/watcher/cmd"
)

var log = logging.MustGetLogger("main")

func main() {
	cmd.Setup()

	exitCode, err := cmd.Parse()
	if err != nil {
		cmd.PrintUsage(err)
		os.Exit(exitCode)
	}

	// NOTE: The spaces here account for the prefixes applied by the logger so that all the text is aligned
	log.Noticef(`%sStarting watcher...
                    Press Ctrl+C or q to quit
                    Press r to to force refresh
 %s`, logging.ColorSeqBold(logging.ColorWhite), []byte("\033[0m"))
	err = watcher.Setup()
	if err != nil {
		log.Fatalf("An error occurred!\nerror: %s", err)
	}
}
