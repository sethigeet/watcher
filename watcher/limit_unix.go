//go:build !windows
// +build !windows

package watcher

import (
	"syscall"

	"github.com/sethigeet/watcher/watcher/cmd"
)

func setupWatchLimit() error {
	var rLimit syscall.Rlimit
	rLimit.Max = *cmd.Config.Limit
	rLimit.Cur = *cmd.Config.Limit

	return syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
}
