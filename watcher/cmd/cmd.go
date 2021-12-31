package cmd

import (
	"flag"
	"fmt"
)

type Config struct {
	Command   *string
	Directory *string
}

var config Config = Config{}

// Setup defines all the allowed flags
func Setup() {
	config.Command = flag.String("cmd", "", "The command to run when a file change is observed")
	flag.StringVar(config.Command, "c", "", "Alias to -cmd")

	config.Directory = flag.String("dir", ".", "The directory to watch for file changes")
	flag.StringVar(config.Directory, "d", ".", "Alias to -dir")
}

// Parse parses all the flags defined by the Setup function
func Parse() (int, error) {
	flag.Parse()

	if len(*config.Command) == 0 {
		return 64, fmt.Errorf("the command must be specified")
	}

	if len(*config.Directory) == 0 {
		return 64, fmt.Errorf("the directory arg must not be empty")
	}

	exists, _ := Exists(*config.Directory)
	if !exists {
		return 66, fmt.Errorf("the directory specified does not exist")
	}

	return 0, nil
}

func PrintUsage(err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		fmt.Println()
	}
	flag.Usage()
}
