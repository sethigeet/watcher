package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mattn/go-zglob"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("main")

func applyIgnoreDefaults(ignoreFileExists bool) error {
	// process the user provided ignore list
	if len(*Config.Ignore) > 0 {
		globsToIgnore := strings.Split(*Config.Ignore, ",")
		for _, glob := range globsToIgnore {
			matches, err := zglob.Glob(*Config.Directory + "/" + glob)
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					continue
				}

				return err
			}

			Config.ToIgnore = append(Config.ToIgnore, matches...)
		}
	}

	// process the user provided ignore file
	if ignoreFileExists {
		file, err := os.Open(*Config.IgnoreFile)
		if err != nil {
			return err
		}
		defer func() {
			if err = file.Close(); err != nil {
				log.Fatal(err)
			}
		}()

		bytes, err := ioutil.ReadAll(file)
		lines := strings.Split(string(bytes), "\n")

		for _, glob := range lines {
			if len(glob) == 0 {
				continue
			}

			matches, err := zglob.Glob(*Config.Directory + "/" + glob)
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					continue
				}

				return err
			}

			Config.ToIgnore = append(Config.ToIgnore, matches...)
		}
	}

	// process the built in default ignore list
	defaultIgnore := []string{
		"**/.git",
		"**/.svn",

		"**/bin",   // a build dir in general
		"**/build", // a build dir in general

		"**/target",       // cargo's(rust pkg manager) build dir
		"**/node_modules", // node.js deps

		"**/*env*", // a virtual environment
	}

	for _, glob := range defaultIgnore {
		matches, err := zglob.Glob(*Config.Directory + "/" + glob)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}

			return err
		}

		Config.ToIgnore = append(Config.ToIgnore, matches...)
	}

	return nil
}

func applyCmdDefaults() error {
	return fmt.Errorf("error not implemented")
}
