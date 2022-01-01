// Package watcher provides functions to watch files in a particular directory
// and run given functions when the files change
package watcher

import (
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/op/go-logging"
	"github.com/sethigeet/watcher/watcher/cmd"
)

var (
	eventsChan chan string
	watchers   []*fsnotify.Watcher
)

func init() {
	eventsChan = make(chan string, 1000)
}

func Setup() error {
	// set the file watchers limit
	setupWatchLimit()

	go handleEvents()

	// setup the watchers
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	go sendEvents(watcher)

	paths := getFilesToWatch()
	for _, path := range paths {
		err = watcher.Add(path)

		if err != nil {
			return err
		}
	}

	if *cmd.Config.ListOnStart {
		log.Infof("%sFiles being watched:%s %s\n", logging.ColorSeqBold(logging.ColorWhite), []byte("\033[0m"), strings.Join(paths, ", "))
	}

	// Close the watcher in the end
	defer watcher.Close()

	// Start by running the cmd
	if *cmd.Config.RunCmdOnStart {
		eventsChan <- *cmd.Config.Directory
	}

	// the quit channel
	quitChan := make(chan os.Signal, 1)

	// accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught
	signal.Notify(quitChan, os.Interrupt)

	// accept 'q' for quitting and 'r' for force refresh
	go detectKeys(quitChan, eventsChan)

	// Block until we receive our signal.
	<-quitChan

	fmt.Printf("\n")
	log.Warning("Disposing of all the watchers...")

	return nil
}
