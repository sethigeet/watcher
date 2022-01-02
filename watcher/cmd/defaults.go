package cmd

import (
	"encoding/json"
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

type defType struct {
	filename string
	apply    func() error
}

var defaults = map[string]defType{
	"nodejs": {
		filename: "package.json",
		apply: func() error {
			content, err := ioutil.ReadFile("./package.json")
			if err != nil {
				return err
			}

			var jsn map[string]interface{}
			err = json.Unmarshal(content, &jsn)
			if err != nil {
				return err
			}

			log.Debugf("jsn: %v", jsn)

			if jsn["scripts"] != nil {
				scripts := (jsn["scripts"]).(map[string]interface{})
				log.Debugf("scripts: %v", scripts)
				if scripts["dev"] != nil {
					*Config.Command = "yarn dev"
					return nil
				}

				if scripts["start"] != nil {
					*Config.Command = "yarn start"
					return nil
				}
			}

			return fmt.Errorf("unable to figure out the default cmd")
		},
	},

	"golang": {
		filename: "go.mod",
		apply: func() error {
			*Config.Command = "go run ."

			return nil
		},
	},

	"rust (managed by cargo)": {
		filename: "Cargo.toml",
		apply: func() error {
			*Config.Command = "cargo run"

			return nil
		},
	},
}

func applyCmdDefaults() error {
	for proj, val := range defaults {
		exists, err := Exists("./" + val.filename)
		if err == nil && exists {
			if val.apply() != nil {
				return err
			}

			log.Noticef("Automatically detected the command for project type: %s", proj)
			return nil
		}
	}

	return fmt.Errorf("this project type is not implemented yet")
}
