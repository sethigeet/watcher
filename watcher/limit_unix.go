//go:build !windows
// +build !windows

package watcher

import (
	"syscall"
)

func setupWatchLimit() error {
	var rLimit syscall.Rlimit
	rLimit.Max = 10000
	rLimit.Cur = 10000

	return syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
}
