package watcher

import (
	"os"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("main")

var format = logging.MustStringFormatter(
	`%{color}[%{time:15:04:05}] ▶ %{color:bold}%{level}%{color:reset} %{message}`,
)

func init() {
	// apply our custom format
	logging.SetFormatter(format)

	backendOut := logging.NewLogBackend(os.Stdout, "", 0)
	backendErr := logging.AddModuleLevel(logging.NewLogBackend(os.Stderr, "", 0))
	backendErr.SetLevel(logging.ERROR, "")

	// Set the backends to be used.
	logging.SetBackend(backendOut, backendErr)
}
