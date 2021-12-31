package cmd

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

type ConfigType struct {
	Command   *string
	Directory *string
	Ignore    *string
	Hidden    *bool

	ToIgnore []string
	RunDelay time.Duration
}

var Config ConfigType = ConfigType{
	RunDelay: 500 * time.Millisecond,
}

// Setup defines all the allowed flags
func Setup() {
	Config.Command = flag.String("cmd", "", "The command to run when a file change is observed")
	flag.StringVar(Config.Command, "c", "", "Alias to -cmd")

	Config.Directory = flag.String("dir", ".", "The directory to watch for file changes")
	flag.StringVar(Config.Directory, "d", ".", "Alias to -dir")

	Config.Ignore = flag.String("ignore", "", "A comma separated list of files to not watch for file changes (supports globbing)")
	flag.StringVar(Config.Ignore, "i", "", "Alias to -ignore")

	Config.Hidden = flag.Bool("hidden", true, "Should the hidden files also be watched for file changes")
	flag.BoolVar(Config.Hidden, "h", true, "Alias to -hidden")
}

// Parse parses all the flags defined by the Setup function
func Parse() (int, error) {
	flag.Parse()

	if len(*Config.Command) == 0 {
		return 64, fmt.Errorf("the command must be specified")
	}

	if len(*Config.Directory) == 0 {
		return 64, fmt.Errorf("the directory arg must not be empty")
	}

	exists, _ := Exists(*Config.Directory)
	if !exists {
		return 66, fmt.Errorf("the directory specified does not exist")
	}

	if len(*Config.Ignore) > 0 {
		globsToIgnore := strings.Split(*Config.Ignore, ",")
		Config.ToIgnore = []string{}
		for _, glob := range globsToIgnore {
			matches, err := Glob(glob)
			if err != nil {
				return 1, err
			}

			Config.ToIgnore = append(Config.ToIgnore, matches...)
		}
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
