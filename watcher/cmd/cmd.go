package cmd

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

type ConfigType struct {
	Command       *string
	Directory     *string
	Ignore        *string
	IgnoreFile    *string
	Hidden        *bool
	Clear         *bool
	RunDelay      *time.Duration
	RunCmdOnStart *bool
	ListOnStart   *bool
	Limit         *uint64

	ToIgnore []string
	CmdParts []string
}

var Config ConfigType = ConfigType{
	ToIgnore: []string{},
}

// Setup defines all the allowed flags
func Setup() {
	Config.Command = flag.String("cmd", "", "The command to run when a file change is observed")
	flag.StringVar(Config.Command, "x", "", "Alias to -cmd")

	Config.Directory = flag.String("dir", ".", "The directory to watch for file changes")
	flag.StringVar(Config.Directory, "d", ".", "Alias to -dir")

	Config.Ignore = flag.String("ignore", "", "A comma separated list of files to not watch for file changes (supports globbing)")
	flag.StringVar(Config.Ignore, "i", "", "Alias to -ignore")
	Config.IgnoreFile = flag.String("ignore-file", "", "A file that contains the names/patterns of the files to not watch(syntax: same as gitignore)")

	Config.Hidden = flag.Bool("hidden", true, "Should the hidden files also be watched for file changes")

	Config.Clear = flag.Bool("clear", false, "Should the hidden files also be watched for file changes")
	flag.BoolVar(Config.Clear, "c", false, "Alias to -clear")

	Config.RunDelay = flag.Duration("delay", 500*time.Millisecond, "The amount of time to wait before running the cmd after a file change occurs")

	Config.RunCmdOnStart = flag.Bool("run-cmd-on-start", true, "Should the specified command run on startup")
	flag.BoolVar(Config.RunCmdOnStart, "r", true, "Alias to -run-cmd-on-start")

	Config.ListOnStart = flag.Bool("list-on-start", false, "Should the files being watched be printed on startup")

	Config.Limit = flag.Uint64("limit", 10000, "The maximum number of files that can be watched")
	flag.Uint64Var(Config.Limit, "l", 10000, "Alias to -limit")
}

// Parse parses all the flags defined by the Setup function
func Parse() (int, error) {
	flag.Parse()

	if len(*Config.Command) == 0 {
		if err := applyCmdDefaults(); err != nil {
			return 64, fmt.Errorf("the command must be specified as it could not be automatically determined")
		}
	}
	Config.CmdParts = strings.Split(*Config.Command, " ")

	if len(*Config.Directory) == 0 {
		return 64, fmt.Errorf("the directory arg must not be empty")
	}

	exists, _ := Exists(*Config.Directory)
	if !exists {
		return 66, fmt.Errorf("the directory specified does not exist")
	}

	if len(*Config.IgnoreFile) > 0 {
		exists, _ = Exists(*Config.IgnoreFile)
		if !exists {
			return 66, fmt.Errorf("the specified ignore file does not exist")
		}
	} else {
		exists = false
	}

	err := applyIgnoreDefaults(exists)
	if err != nil {
		return 1, err
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
