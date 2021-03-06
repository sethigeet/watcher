package watcher

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/mattn/go-tty"
	"github.com/sethigeet/watcher/watcher/cmd"
)

func isIgnored(name string) bool {
	for _, file := range cmd.Config.ToIgnore {
		if name == file {
			return true
		}
	}
	return false
}

func getFilesToWatch() []string {
	files := []string{}

	filepath.WalkDir(*cmd.Config.Directory, func(path string, d fs.DirEntry, err error) error {
		// Skip the file/dir if it is to be ignored
		if isIgnored(path) {
			return filepath.SkipDir
		}

		// Skip the file/dir if it is hidden and the config says so
		if !*cmd.Config.Hidden {
			if len(path) > 1 && strings.HasPrefix(filepath.Base(path), ".") {
				return filepath.SkipDir
			}
		}

		// setup the watcher if the "path" is not a directory
		if !d.IsDir() {
			files = append(files, path)
			return nil
		}

		return err
	})

	return files
}

func detectKeys(quitChan chan<- os.Signal, refreshChan chan<- string) {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {
		input, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		switch input {
		case 'q':
			quitChan <- os.Interrupt
		case 'r':
			refreshChan <- *cmd.Config.Directory
		}
	}
}
